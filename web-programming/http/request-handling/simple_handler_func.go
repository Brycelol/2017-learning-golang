// This demonstrates handler functions
// Handler functions behave similarly to handlers in that they have the same
// signatures as ServeHTTP.
//
// These work because go has a function type named HandlerFunc. All functions f
// matching this HandlerFunc signature will be adapted into a handler with a method f
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World")
}

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// Handle func takes our handler and calls HandlerFunc(f)
	// underneath to convert the function into a handler.
	// eg helloHandler := HandlerFunc(hello)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()

	// Handler functions might seem cleaner than using handlers - doing the job just
	// as well. It all boils down to design.
	// If you have an existing interface or want a type that can be used as a handler - simply
	// add a ServeHTTP method. This can allow you to build more modular web applications (although
	// remember the SRP! *but Go isn't object orientated)
}
