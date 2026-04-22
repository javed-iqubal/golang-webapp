package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world ! Welcome to Golang learning page")
	})

	// start web server

	server := http.Server{
		Addr:    ":8082",
		Handler: nil,
	}

	fmt.Println("Server is running...")
	server.ListenAndServe()
}
