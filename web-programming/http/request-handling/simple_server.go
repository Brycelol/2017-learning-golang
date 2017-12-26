package main

import (
	"net/http"
)

func main() {

	// With no address specified this listens on 127.0.0.1:80
	// Because the handler parameter is nil, the default multiplexer, DefaultServeMux, is used
	http.ListenAndServe("", nil)
}
