package main

import (
	"fmt"
	"net/http"
)

// HelloHandler struct that we will bind the ServerHTTP handler interface to
type HelloHandler struct{}

// Interface method for handlers - Bound to the HelloHandler struct
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// WorldHandler struct that we will bind the ServerHTTP handler interface to
type WorldHandler struct{}

// Interface method for handlers - Bound to the HelloHandler struct
func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World")
}

func main() {

	// Create our two handlers, both bound with ServeHTTP implementations for handlers
	helloHandler := HelloHandler{}
	worldHandler := WorldHandler{}

	// Handler is left nil, so we will use the DefaultServeMux multiplexer
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// http.Handle attaches a handler to the DefaultServeMux multiplexer
	// The Handle function is also a method of DefaultServeMux - when you call
	// http.Handle you are actually calling the DefaultServeMux's Handle method
	// it's exposed in http.Handle for convenience
	http.Handle("/hello", &helloHandler)
	http.Handle("/world", &worldHandler)

	server.ListenAndServe()
}
