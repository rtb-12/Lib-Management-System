package models

import (
	"fmt"
)

func AddBook(bookName string){
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	insertSql := "INSERT INTO books_list (name) VALUES (?)"
	_, err = db.Exec(insertSql, bookName)
	if err != nil {
		fmt.Printf("error %s inserting into the database", err)
	} else {
		fmt.Printf("successfully inserted %s into the database", bookName)
	}
}


