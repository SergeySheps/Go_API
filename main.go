package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const port int = 3000

var books []Book

//Book struct
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, book := range books {
		if params["id"] == book.ID {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))

	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func editBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for ind, item := range books {
		if book.ID == item.ID {
			books = append(books[:ind], books[ind+1:]...)
			books = append(books, book)
			break
		}
	}

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for ind, item := range books {
		if params["id"] == item.ID {
			books = append(books[:ind], books[ind+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {

	book := Book{ID: "1", Title: "Book one", Author: &Author{FirstName: "John", LastName: "Konor"}}
	books = append(books, book)

	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", editBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Printf("Server started at port: %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
