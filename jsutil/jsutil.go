// Package jsutil provides various convenience functions for Javascript environment
package jsutil

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/robertkrimen/otto"
)

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
