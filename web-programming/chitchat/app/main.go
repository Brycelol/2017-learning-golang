// Main package for the hello-world go web server
package main

import (
	"net/http"
	"github.com/brycelol/learning-golang/web-programming/chitchat/app/data"
)

// This will be our index handler
// Default entry point for the web app
func index(w http.ResponseWriter, r *http.Request) {

	// Getting a list of the current threads on the chat app
	threads, err := data.Threads()

	if err == nil {

		// Authenticate the session using out utility function
		// an error returned here tells us that the session is invalid
		// session returns a Session struct but we aren't interested in that
		// right now so we assign it to _
		_, err := session(w, r)

		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}

}

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
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
