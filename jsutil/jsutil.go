// Package jsutil provides various convenience functions for Javascript environment
package jsutil

import (
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

func ConfigureRequire(vm *otto.Otto) {
	vm.Set("require", func(call otto.FunctionCall) otto.Value {

		const requireHeader = `(function() {
    var module = { exports: {} };
    var exports = module.exports;`

		const requireFooter = `   return module.exports;
}());`
		path, err := call.Argument(0).ToString()
		if err != nil {
			return otto.UndefinedValue()
		}

		body, err := ioutil.ReadFile(path)
		if err != nil {
			return otto.UndefinedValue()
		}

		val, err := vm.Run(requireHeader + string(body) + requireFooter)
		if err != nil {
			return otto.UndefinedValue()
		}

		return val
	})
}
