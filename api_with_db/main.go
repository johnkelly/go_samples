package main

import (
	"log"
	"net/http"

	"github.com/johnkelly/go_samples/api_with_db/books"
)

func main() {
	http.HandleFunc("/books", dbMiddleware(books.IndexHandler, db))
	http.HandleFunc("/books/show", dbMiddleware(books.ShowHandler, db))
	http.HandleFunc("/books/create", dbMiddleware(books.CreateHandler, db))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
