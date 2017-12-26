package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// Use makecert.go to generate the pub / private key pair
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
