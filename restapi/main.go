package main

import (
<<<<<<< HEAD
	"log"
	"net/http"

	"github.com/gorilla/mux"
=======
	// O pacote json implementa a codificação e a decodificação de JSON
	"log"      //O log de pacotes implementa um pacote de log simples. Ele define um tipo, Logger, com métodos para formatação de saída.
	"net/http" // O pacote http fornece implementações de cliente e servidor HTTP. Get, Head, Post e PostForm fazem solicitações HTTP (ou HTTPS)
	// Pacote aritmético/ rand gera número aleatório
	// O pacote strconv implementa conversões de x para representações de strings de tipos de dados básicos.

	"github.com/gorilla/mux" // Dependência do MUX
>>>>>>> a18ed3d4682c9bbf4fa418ab133001d88488439b
)

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

<<<<<<< HEAD
// Init Books var as a slice Book struct
var books []Book

=======
>>>>>>> a18ed3d4682c9bbf4fa418ab133001d88488439b
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

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "847564", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
