package main

import (
	"fmt"
	"net/http"
)

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to appointment app called mementoludi!</h1>")
}

func contactsHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contacts page</h1>")
}

func pageNotFoundHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func pathHandle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		mainPageHandle(w, r)
	case "/contacts":
		contactsHandle(w, r)
	default:
		// Handle the page not found error
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandle)
	fmt.Println("Starting server on localhost:3000...")
	fmt.Println("Press Ctrl + C to stop the server")
	http.ListenAndServe(":3000", nil)
}
