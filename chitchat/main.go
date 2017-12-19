// Main package for the hello-world go web server
package main

import (
	"net/http"
)

func main() {

	// This is our multiplexer
	// the multiplexer will inspect a URL being requested
	// and redirect it to the correct handler
	mux := http.NewServeMux()

	// FileServer serves files from a specified directory
	files := http.FileServer(http.Dir("/public"))

	// Note here: Handlers & Handler funcs are NOT the same (although same results
	// in the end).
	// Note StripPrefix - Used to strip the given prefix from the request URL's path

	// We are basically saying here: all requests going to /static/, strip the /static
	// and look for the files in /public
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Adding a handler to the multiplexer
	// first parameter is the route, second is the handler
	// handler will always take a http response / request
	//mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
