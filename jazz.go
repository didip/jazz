package jazz

import (
	"fmt"
	"net/http"

	"github.com/robertkrimen/otto"

	"github.com/didip/jazz/httputil"
	"github.com/didip/jazz/jsutil"
)

// JSFuncHandler loads and runs a javascript module that can process http request given request handler function.
func JSFuncHandler(vm *otto.Otto, jsModule string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsutil.ConfigureRequire(vm)
		vm.Set("request", r)
		vm.Set("response", w)
		vm.Set("ResponseUtil", httputil.NewResponseUtil)

		_, err := vm.Run(fmt.Sprintf(`require('%v').handle(request, response);`, jsModule))
		if err != nil {
			w.Header().Add("Content-Type", "text/html") // TODO(didip): This has to be configurable.
			w.WriteHeader(500)                          // TODO(didip): This has to be configurable.
			w.Write([]byte(err.Error()))
			return
		}
	}
}
