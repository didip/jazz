[![GoDoc](https://godoc.org/github.com/didip/jazz?status.svg)](http://godoc.org/github.com/didip/jazz)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/didip/jazz/master/LICENSE.md)

Handle HTTP requests with style (in Javascript).


## Motivation

> Javascript and Go belong together. I should be able to use npm modules inside my Go runtime.


## Design Goals

1. Simple and consistent handler API on JS land.

2. Able to require npm modules. Must handle npm load order correctly.

3. Profit!


## Five Minutes Tutorial

```
// 1. Write your javascript modules. Example: jshandlers/index.js
// module.exports = {
// 	 handle: function(request, response) {
//	 	 ResponseUtil(response).WriteString("Hello World, from Javascript")
//	 }
// }

// 2. Write your Go project
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

	// Configure all the global settings
	jsutil.Configure(vm)

	// Expected output: "Hello World, from Go"
	http.HandleFunc("/", gohello)

	// Expected output: "Hello World, from Javascript"
	http.HandleFunc("/js", jazz.JSFuncHandler(vm, "jshandlers/index.js"))

	http.ListenAndServe(":8080", nil)
}
```


## Why do I need this?

1. You want to provide extensibility to your Go HTTP service.

2. You have this very important npm module that does not have Go equivalent.

3. You are just hacking for fun, wanting to see what's possible.


## Otto Caveats

* It complies to ECMAScript 5.1.

* "use strict" will parse, but does nothing.

* The regular expression engine (re2/regexp) is [not fully compatible](https://github.com/robertkrimen/otto#regular-expression-incompatibility) with the ECMA5 specification.
