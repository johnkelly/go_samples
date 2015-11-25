package main

import (
	"log"
	"net/http"

	"github.com/johnkelly/go_samples/api_with_db/book"
)

func main() {
	http.HandleFunc("/books", dbMiddleware(book.IndexHandler, db))
	http.HandleFunc("/books/show", dbMiddleware(book.ShowHandler, db))
	http.HandleFunc("/books/create", dbMiddleware(book.CreateHandler, db))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
