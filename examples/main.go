package main

import (
	"fmt"
	"net/http"

	"github.com/didip/jazz"
	"github.com/robertkrimen/otto"
)

func gohello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, from Go")
}

func main() {
	vm := otto.New()

	http.HandleFunc("/", gohello)
	http.HandleFunc("/js", jazz.JSFuncHandler(vm, "jshandlers/index.js"))

	http.ListenAndServe(":8080", nil)
}
