package models

import (
	"fmt"

	"github.com/ashpect/go-mvc/pkg/types"
)

func FetchBooks() types.ListBooks {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM books_list"
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Name)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchBooks = append(fetchBooks, book)
	}

	var listBooks types.ListBooks
	listBooks.Books = fetchBooks
	return listBooks

}
