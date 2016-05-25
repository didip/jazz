module.exports = {
	handle: function(request, response) {
		ResponseWriteString(response, "Hello World, from Javascript")
	}
}
