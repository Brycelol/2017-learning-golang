package main

import "net/http"

func main() {

	// A Server defines parameters for running an HTTP server.
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// Listen and serve using the our defined server
	server.ListenAndServe()
}
