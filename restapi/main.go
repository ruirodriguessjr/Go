package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com\gorilla\mux"
)

// Author Struct
type Author struct {
	Fistname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	
}

//Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	// Init Route
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
