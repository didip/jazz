# Jazz

Handle HTTP requests with style! (in Javascript)


## Motivation

> Javascript and Go belong together. I should be able to use NPM modules for handling requests!


## Basic Concepts

1. Simple handler API on JS land. Ideas:
	```javascript
	// /jshandlers/index.js
	module.exports = {
		handle: function(request, response) {
			ResponseWriteString(response, "Hello World, from Javascript")
		}
	}
	```

2. Able to require NPM modules.

3. Profit!
