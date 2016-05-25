[![GoDoc](https://godoc.org/github.com/didip/jazz?status.svg)](http://godoc.org/github.com/didip/jazz)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/didip/jazz/master/LICENSE.md)

Handle HTTP requests with style (in Javascript).


## Motivation

> Javascript and Go belong together. I should be able to use npm modules inside my Go runtime.


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


## Features

1. Write request handlers in Javascript.

2. Javascript handlers are loaded during runtime, so you can refresh and see changes immediately. Just like PHP.

3. Able to require npm modules. `$NODE_PATH` is used to find modules.

4. The javascript VM can further be enhanced to provide more functionalities.


## Why do I need this?

1. You want to provide extensibility to your Go HTTP service.

2. You have this very important npm module that does not have Go equivalent.

3. You are just hacking for fun, wanting to see what's possible.


## Otto Caveats

* It complies to ECMAScript 5.1.

* "use strict" will parse, but does nothing.

* The regular expression engine (re2/regexp) is [not fully compatible](https://github.com/robertkrimen/otto#regular-expression-incompatibility) with the ECMA5 specification.
