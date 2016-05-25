package main

import (
	"fmt"
	"net/http"

	"github.com/didip/jazz"
	"github.com/didip/jazz/jsutil"
	"github.com/robertkrimen/otto"
)

func gohello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, from Go")
}

func main() {
	vm := otto.New()

	// Configure all the global settings
	jsutil.Configure(vm)

	// Expected output: "Hello World, from Go"
	http.HandleFunc("/", gohello)

	// Expected output: "Hello World, from Javascript"
	http.HandleFunc("/js", jazz.JSFuncHandler(vm, "jshandlers/index.js"))

	http.ListenAndServe(":8080", nil)
}
