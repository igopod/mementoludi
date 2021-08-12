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

func main() {
	http.HandleFunc("/", mainPageHandle)
	http.HandleFunc("/contacts", contactsHandle)
	fmt.Println("Starting server on localhost:3000...")
	fmt.Println("Press Ctrl + C to stop the server")
	http.ListenAndServe(":3000", nil)
}
