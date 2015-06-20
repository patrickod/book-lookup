package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func recordBook(isbn string, book BookInfo) {
	bookid := insertBook(isbn, book)
	authorids := insertAuthors(bookid, book)
	insertAuthorships(bookid, authorids)
}

func insertBook(isbn string, book BookInfo) int {
	db, err := sql.Open("postgres", "postgresql://localhost:5432/noisebridge_library?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	var bookid int
	err = db.QueryRow(`INSERT INTO books(isbn, title, number_of_pages, cover_image, publish_date) VALUES($1, $2, $3, $4, $5) RETURNING id`, isbn, book.Title, book.NumberOfPages, book.Cover.Large, nil).Scan(&bookid)
	if err != nil {
		log.Fatal(err)
	}
	return bookid
}

func insertAuthors(bookid int, book BookInfo) []int {
	var authorids []int
	for _, author := range book.Authors {
		authorid := insertAuthor(author)
		authorids = append(authorids, authorid)
	}
	return authorids
}

func insertAuthor(author Author) int {
	db, err := sql.Open("postgres", "postgresql://localhost:5432/noisebridge_library?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	var authorid int
	err = db.QueryRow(`INSERT INTO authors(name, url) VALUES($1, $2) RETURNING id`, author.Name, author.Url).Scan(&authorid)
	if err != nil {
		log.Fatal(err)
	}
	return authorid
}

func insertAuthorships(bookid int, authorids []int) {
	db, err := sql.Open("postgres", "postgresql://localhost:5432/noisebridge_library?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	for _, authorid := range authorids {
		_, err = db.Exec(`INSERT INTO authorships(author_id, book_id) VALUES($1, $2)`, authorid, bookid)
		if err != nil {
			log.Fatal(err)
		}
	}
}
