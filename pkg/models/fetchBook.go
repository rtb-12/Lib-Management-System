package models

import (
	"fmt"

	"github.com/rtb-12/Lib-Management-System/pkg/types"
)

func FetchBooks() types.ListBooks {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM book_info WHERE QuantityAvailable>0"
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchBooks []types.BookInfo
	for rows.Next() {
		var book types.BookInfo
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.QuantityAvailable)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchBooks = append(fetchBooks, book)
	}

	var listBooks types.ListBooks
	listBooks.Books = fetchBooks
	return listBooks

}
