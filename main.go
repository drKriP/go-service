package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	books := []Book{
		{ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
		{ID: "2", Title: "Clean Code", Author: "Robert C. Martin"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Microservice!")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/books", getBooks)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
