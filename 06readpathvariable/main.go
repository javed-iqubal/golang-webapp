package main

import (
	"fmt"
	"net/http"
)

func main() {
	/**
	// Old way to read path parameters
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {

		prefix := "/user/"
		if len(r.URL.Path) <= len(prefix) {
			http.Error(w, "User id is required (e.g. /user/1234)", http.StatusBadRequest)
		}

		id := r.URL.Path[len(prefix):]
		id = strings.Split(id, "/")[0]

		fmt.Fprintf(w, "User Id is : %s", id)
	})
	*/
	// Latest way to read path parameters

	http.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		fmt.Fprintf(w, "User id is : %s", id)
	})

	fmt.Println("Server is running.....")
	http.ListenAndServe(":8082", nil)
}
