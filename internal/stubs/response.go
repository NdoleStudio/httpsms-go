package stubs

// HttpInternalServerErrorResponse internal error response
func HttpInternalServerErrorResponse() []byte {
	return []byte(`
		{
			"message": "We ran into an internal error while handling the request.",
			"status": "error"
		}
`)
}
