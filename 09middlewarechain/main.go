package main

import (
	"fmt"
	"net/http"
	"time"
)

// loggin middleware
func LogginMiddleware(next http.Handler) http.Handler {

	// hof
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// logic before the handler
		fmt.Printf("Started: %s | %s | %s \n", r.RemoteAddr, r.Method, r.URL.Path)

		// call the next handler in chain
		w.Header().Add("companyName", "HCL Technologies Ltd")
		next.ServeHTTP(w, r)

		// logic after the handler
		fmt.Fprintf(w, "Completed in %s", time.Since(start))
	})
}

// Auth middlware : check for a specific header

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("X-Auth-Token")
		if token != "secret-pass" {
			fmt.Println("[Auth] Unauthorized Attempt")
			http.Error(w, "Forbidden: Invalid Token", http.StatusForbidden)
			return // stop the chain here
		}
		fmt.Println("[Auth] User Authenticated")
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		// logic
		fmt.Fprintln(w, "Hello, Javed")

	})

	// wrap the mux with middleware

	// order -> request ---> logging ---> Auth--Mux--Handler

	// middleware chain a(b(c(d(e(mux)))))

	wrappedMux := LogginMiddleware(AuthMiddleware(mux))

	fmt.Println("Server is running.....")
	err := http.ListenAndServe(":8082", wrappedMux)
	if err != nil {
		fmt.Printf("Error starting server: %s", err.Error())
	}
}
