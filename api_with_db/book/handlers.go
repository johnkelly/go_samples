package book

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

//ShowHandler is the book show action
func ShowHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		w.WriteHeader(400)
		return
	}

	book, err := findBookByISBN(db, isbn)

	if err == sql.ErrNoRows {
		w.WriteHeader(404)
		return
	} else if err != nil {
		w.WriteHeader(500)
		return
	}

	if err := json.NewEncoder(w).Encode(book); err != nil {
		w.WriteHeader(500)
		return
	}
}

//IndexHandler is the book index action
func IndexHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	books, err := allBooks(db)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		w.WriteHeader(500)
		return
	}
}

//CreateHandler is the book create action
func CreateHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	isbn := r.FormValue("isbn")
	title := r.FormValue("title")
	author := r.FormValue("author")

	if isbn == "" || title == "" || author == "" {
		w.WriteHeader(400)
		return
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = createBook(db, isbn, title, author, price)

	if err != nil {
		w.WriteHeader(422)
		return
	}

	w.WriteHeader(201)
}
