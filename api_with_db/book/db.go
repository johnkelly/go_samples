package book

import (
	"database/sql"
	"log"
)

//Book is a struct that contains isbn, title, author, & price.
type Book struct {
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func allBooks(db *sql.DB) ([]*Book, error) {
	rows, err := db.Query("SELECT * FROM books")

	var books []*Book

	if err != nil {
		return books, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return books, err
	}
	return books, nil
}

func findBookByISBN(db *sql.DB, isbn string) (*Book, error) {
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)
	book := new(Book)
	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	if err != nil {
		return book, err
	}
	return book, nil
}

func createBook(db *sql.DB, isbn string, title string, author string, price float64) error {
	_, err := db.Exec("INSERT INTO books VALUES($1, $2, $3, $4)", isbn, title, author, price)

	if err != nil {
		return err
	}

	return nil
}
