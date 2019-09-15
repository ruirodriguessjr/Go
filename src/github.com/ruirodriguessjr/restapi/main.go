package main

import (
	"context"
	"encoding/json" // O pacote json implementa a codificação e a decodificação de JSON
	"fmt"
	"log"      //O log de pacotes implementa um pacote de log simples. Ele define um tipo, Logger, com métodos para formatação de saída.
	"net/http" // O pacote http fornece implementações de cliente e servidor HTTP. Get, Head, Post e PostForm fazem solicitações HTTP (ou HTTPS)

	"math/rand" // Pacote aritmético/ rand gera número aleatório
	"strconv"   // O pacote strconv implementa conversões de x para representações de strings de tipos de dados básicos.

	"github.com/gorilla/mux" // Dependência do MUX
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init Books var book and as a slice Books struct

var books []Book
var book Book

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Content-Type faz com que ele retorne para mim
	//qual o formato de retorno que eu espero
	w.Header().Set("Content-Type", "application/json")
	// NewEncoder e Encode eu estou escrevendo o meu response
	// Dentro do meu objeto passando pelo formato Json
	json.NewEncoder(w).Encode(books)
}

// Get single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode((&book))
	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID = not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode((&book))
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Init Route
	router := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "447539", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "843921", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "123654", Title: "Book Tree", Author: &Author{Firstname: "Michael", Lastname: "Jackson"}})
	books = append(books, Book{ID: "4", Isbn: "654532", Title: "Book Four", Author: &Author{Firstname: "James", Lastname: "Hetfield"}})

	// Route Handlers / Endpoints(tipo url) - Função Handler manipula aquela requisição
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// ListenAndServer é um servidor padrão do Go,
	//onde eu passo a rota passando a sua porta de acesso,
	// Todas as linguagens tem isso, pelo menos a maioria.
	log.Fatal(http.ListenAndServe(":8000", router))
	// log.Fatal caso ocorra um erro, ele aborta toda a minha conexão

}
