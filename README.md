# Jazz

Handle HTTP requests with style! (in Javascript)


## Motivation

> Javascript and Go belong together. I should be able to use NPM modules for handling requests!


## Basic Concepts

1. Simple handler API on JS land. Ideas:
	```javascript
	// Use the response struct from Go.
	// /jshandlers/index.js
	module.exports = {
		handle: function(request, response) {
			response.Write("<html><head>My JS Page</head><body>Hello World, from JS</body></html>")
		}
	}
	```

2. Able to require NPM modules.

3. Profit!
