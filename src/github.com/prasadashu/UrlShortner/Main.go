package main

import (
	"fmt"
	"net/http"
)

var endpointMap = map[string]string{
	"/linkedin": "https://www.linkedin.com/in/ashuprasad/",
	"/github":   "https://github.com/prasadashu/",
}

func endpointHandler(write http.ResponseWriter, request *http.Request) {
	// Get URL path from request
	requestPath := request.URL.Path

	// Check if URL exist in endpoint map
	if urlPath, ok := endpointMap[requestPath]; ok {
		// Redirect user to path
		http.Redirect(write, request, urlPath, http.StatusFound)
		return
	}

	// Otherwise, URL does not exist
	// http.NotFound(write, request)

	// Custom 404 response
	write.WriteHeader(http.StatusNotFound)
	write.Write([]byte("404: Resource not found\n"))
}

func main() {
	// Run server on port 9000
	fmt.Println("Server running on port 9000")
	http.ListenAndServe(":9000", http.HandlerFunc(endpointHandler))
}
