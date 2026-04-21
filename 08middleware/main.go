package main

import (
	"fmt"
	"net/http"
	"time"
)

// middleware
func loggingMiddleware(next http.Handler) http.Handler {

	// hof
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// logic before the handler
		fmt.Printf("Started: %s | %s | %s \n", r.RemoteAddr, r.Method, r.URL.Path)

		// call the next handler in chain
		next.ServeHTTP(w, r)

		// logic after the handler
		w.Header().Add("companyName", "HCL Technologies Ltd")
		fmt.Fprintf(w, "Completed in %s", time.Since(start))
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		// logic
		fmt.Fprintln(w, "Hello, Javed")

	})

	// wrap the mux with middleware

	wrappedMux := loggingMiddleware(mux)

	fmt.Println("Server is running.....")
	err := http.ListenAndServe(":8082", wrappedMux)
	if err != nil {
		fmt.Printf("Error starting server: %s", err.Error())
	}
}
