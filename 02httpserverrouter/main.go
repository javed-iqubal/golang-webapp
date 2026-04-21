package main

import (
	"fmt"
	"net/http"
)

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	// response.Write([]byte("Welcome to my home page"))
	// fmt.Fprintf(response, "<h1>My Home Page</h1>")
	fmt.Fprintln(response, "Welcome to my home page")
}

func main() {
	// router - application logic

	// http.HandleFunc("/", HomePageHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world ! Welcome to home page")
	})

	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
