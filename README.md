[![GoDoc](https://godoc.org/github.com/didip/jazz?status.svg)](http://godoc.org/github.com/didip/jazz)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/didip/jazz/master/LICENSE.md)

# Jazz

Handle HTTP requests with style! (in Javascript)


## Motivation

> Javascript and Go belong together. I should be able to use NPM modules for handling requests!


## Design Goals

1. Simple and consistent handler API on JS land.

2. Able to require NPM modules. Must handle NPM load order correctly.

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

	http.HandleFunc("/", gohello)
	http.HandleFunc("/js", jazz.JSFuncHandler(vm, "jshandlers/index.js"))

	http.ListenAndServe(":8080", nil)
}
```