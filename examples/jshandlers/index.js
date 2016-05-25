module.exports = {
	handle: function(request, response) {
		ResponseUtil(response).WriteString("Hello World, from Javascript")
	}
}
