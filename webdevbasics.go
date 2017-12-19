package main

import (
	"fmt"
	"net/http")

func wdIndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,`
		<h1>Hello there.</h1>
		<p>Go is fast...</p>
		<p>We can template values [%s]</p>
	`, "<strong>VAL</strong>")

	// HTML templates would be better here...just a demo.
}

func main() {
	http.HandleFunc("/", wdIndexHandler)
	http.ListenAndServe(":8000", nil)
}
