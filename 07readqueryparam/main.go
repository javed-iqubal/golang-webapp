package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {

		queryParam := r.URL.Query()
		name := queryParam.Get("name")
		if name == "" {
			name = "guest"
		}

		fmt.Fprintln(w, "Hello", name)
	})

	fmt.Println("Server is running....")
	http.ListenAndServe(":8082", nil)
}
