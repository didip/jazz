// Package jsutil provides various convenience functions for Javascript environment
package jsutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/didip/jazz/httputil"
	"github.com/robertkrimen/otto"
)

type jstimer struct {
	timer    *time.Timer
	duration time.Duration
	interval bool
	call     otto.FunctionCall
}

// Configure enhances the JS VM.
func Configure(vm *otto.Otto) {
	ConfigureRequire(vm)
	ConfigureTimeoutInterval(vm)

	vm.Set("ResponseUtil", httputil.NewResponseUtil)
}

// ConfigureRequire provides require() functionality to JS VM.
func ConfigureRequire(vm *otto.Otto) {
	vm.Set("require", func(call otto.FunctionCall) otto.Value {

		const requireHeader = `(function() {
    var module = { exports: {} };
    var exports = module.exports;`

		const requireFooter = `   return module.exports;
}());`
		modulePath, err := call.Argument(0).ToString()
		if err != nil {
			return otto.UndefinedValue()
		}

		ableToReadModule := false

		// Check current directory to load module
		body, err := ioutil.ReadFile(modulePath)
		if err == nil {
			ableToReadModule = true
		}

		if !ableToReadModule {
			// Check $NODE_PATH to load module
			nodePaths := os.Getenv("NODE_PATH")
			if nodePaths != "" {
				for _, nodePath := range filepath.SplitList(nodePaths) {
					body, err = ioutil.ReadFile(filepath.Join(nodePath, modulePath))
					if err == nil {
						ableToReadModule = true
						break
					}
				}
			}
		}

		if !ableToReadModule {
			return otto.UndefinedValue()
		}

		val, err := vm.Run(requireHeader + string(body) + requireFooter)
		if err != nil {
			return otto.UndefinedValue()
		}

		return val
	})
}

// ConfigureTimeoutInterval provides setTimeout() and setInterval() functionalities to JS VM.
func ConfigureTimeoutInterval(vm *otto.Otto) {
	registry := map[*jstimer]*jstimer{}
	ready := make(chan *jstimer)

	newTimer := func(call otto.FunctionCall, interval bool) (*jstimer, otto.Value) {
		delay, _ := call.Argument(1).ToInteger()
		if 0 >= delay {
			delay = 1
		}

		timer := &jstimer{
			duration: time.Duration(delay) * time.Millisecond,
			call:     call,
			interval: interval,
		}
		registry[timer] = timer

		timer.timer = time.AfterFunc(timer.duration, func() {
			ready <- timer
		})

		value, err := call.Otto.ToValue(timer)
		if err != nil {
			return timer, otto.UndefinedValue()
		}

		return timer, value
	}

	setTimeout := func(call otto.FunctionCall) otto.Value {
		_, value := newTimer(call, false)
		return value
	}
	vm.Set("setTimeout", setTimeout)

	setInterval := func(call otto.FunctionCall) otto.Value {
		_, value := newTimer(call, true)
		return value
	}
	vm.Set("setInterval", setInterval)

	clearTimeout := func(call otto.FunctionCall) otto.Value {
		timer, _ := call.Argument(0).Export()
		if timer, ok := timer.(*jstimer); ok {
			timer.timer.Stop()
			delete(registry, timer)
		}
		return otto.UndefinedValue()
	}
	vm.Set("clearTimeout", clearTimeout)
	vm.Set("clearInterval", clearTimeout)

	go func() {
		for {
			select {
			case timer := <-ready:
				var arguments []interface{}
				if len(timer.call.ArgumentList) > 2 {
					tmp := timer.call.ArgumentList[2:]
					arguments = make([]interface{}, 2+len(tmp))
					for i, value := range tmp {
						arguments[i+2] = value
					}

				} else {
					arguments = make([]interface{}, 1)
				}

				arguments[0] = timer.call.ArgumentList[0]

				_, err := vm.Call(`Function.call.call`, nil, arguments...)
				if err != nil {
					for _, timer := range registry {
						timer.timer.Stop()
						delete(registry, timer)
					}
				}

				if timer.interval {
					timer.timer.Reset(timer.duration)
				} else {
					delete(registry, timer)
				}

			default:
				// Escape valve!
				// If this isn't here, we deadlock...
			}

			if len(registry) == 0 {
				break
			}
		}
	}()
}
