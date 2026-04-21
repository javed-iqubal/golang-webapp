package main

import (
	"fmt"
	"net/http"
)

func main() {

	// router-1: hello page

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})
	// router-2: home page
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to home page!")
	})

	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
