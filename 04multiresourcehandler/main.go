package main

import (
	"fmt"
	"net/http"
)

// User resource
func UserHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "User Get.")
	case http.MethodPost:
		fmt.Fprintf(w, "User Post.")
	case http.MethodPut:
		fmt.Fprintf(w, "User Put.")
	case http.MethodDelete:
		fmt.Fprintf(w, "User Delete.")

	}
}

// Customer resource
func CustomerHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Customer Get.")
	case http.MethodPost:
		fmt.Fprintf(w, "Customer Post.")
	case http.MethodPut:
		fmt.Fprintf(w, "Customer Put.")
	case http.MethodDelete:
		fmt.Fprintf(w, "Customer Delete.")

	}
}

// Product resource
func ProductHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Product Get.")
	case http.MethodPost:
		fmt.Fprintf(w, "Product Post.")
	case http.MethodPut:
		fmt.Fprintf(w, "Product Put.")
	case http.MethodDelete:
		fmt.Fprintf(w, "Product Delete.")

	}
}

func main() {

	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/customer", CustomerHandler)
	http.HandleFunc("/product", ProductHandler)

	fmt.Println("Server is running...")
	http.ListenAndServe(":8082", nil)
}
