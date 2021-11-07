package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Init books var as a slice Book struct

var books []Book

// Get all Books

func getBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&books)
}

func getBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // GET params
	// Loop through books and find with ID

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(rw).Encode(item)
			return
		}
	}
	json.NewEncoder(rw).Encode(&Book{})

}

func createBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var book Book

	// Important: below operation decodes the request body and setting to the book struct
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(rw).Encode(&book)

}

func updateBook(rw http.ResponseWriter, r *http.Request) {

}

func deteleBook(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	fmt.Println("Hello World")
	// Init Router
	r := mux.NewRouter()

	books = append(books, Book{
		ID:    "1",
		Isbn:  "12345",
		Title: "Book one",
		Author: &Author{
			FirstName: "Akshay",
			LastName:  "Iyyadurai Balasundaram",
		},
	})

	books = append(books, Book{
		ID:    "2",
		Isbn:  "98765",
		Title: "Book Two",
		Author: &Author{
			FirstName: "Becky",
			LastName:  "Iyyadurai Balasundaram",
		},
	})

	// Route Handles / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods(http.MethodGet)
	r.HandleFunc("/api/books/{id}", getBook).Methods(http.MethodGet)
	r.HandleFunc("/api/books", createBook).Methods(http.MethodPost)
	r.HandleFunc("/api/books/{id}", updateBook).Methods(http.MethodPut)
	r.HandleFunc("/api/books/{id}", deteleBook).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8000", r))

}
