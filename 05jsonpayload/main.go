package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		// read and write

		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// response

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	})

	fmt.Println("Server is running....")
	http.ListenAndServe(":8082", nil)
}
