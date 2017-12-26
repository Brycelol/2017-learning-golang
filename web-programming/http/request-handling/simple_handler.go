// Demonstrates a simple handler in go
package main

import (
	"fmt"
	"net/http"
)

// Empty struct representing our handler
type HelloWorldHandler struct{}

// Method attached to a pointer to the HelloWorldHandler struct.
//
// In Go, a handler is an interface with a method named ServeHTTP
// which takes a response writer and pointer to a request a parameter.
//
// Essentially, anything that has a method ServeHTTP matching this signature
// is a handler
func (h *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	// Create our simple hello world handler
	helloWorldHandler := HelloWorldHandler{}

	// We pass through a pointer to our handler here, which will handle all
	// requests with a hello world response.
	//
	// Note that when Handler is nil, it defaults to DefaultServeMux
	// which is a multiplexer. This is because ServeMux (which DefaultServeMux is an instance of)
	// has a ServeHTTP with the same signature - hence it is also an instance of the handler struct.
	// It's a special handler though, as its job is to redirect to different handlers depending on the URL.
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &helloWorldHandler,
	}

	server.ListenAndServe()
}
