package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/rtb-12/Lib-Management-System/pkg/types"
)

func checkBookExists(book types.BookInfo) bool {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	var name string
	// Adjust the SQL query to check both title and author
	err = db.QueryRow("SELECT Title FROM book_info WHERE Title = ? AND Author = ?", book.Title, book.Author).Scan(&name)
	if err != nil {
		return false
	}
	return true
}

func AddBook(book types.BookInfo) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()

	// Assuming checkBookExists is modified to return (bool, error)
	isBookExists := checkBookExists(book)

	if !isBookExists {
		insertSql := "INSERT INTO book_info (Title, Author, Genre, Stock) VALUES (?, ?, ?, ?)"
		_, err = db.Exec(insertSql, book.Title, book.Author, book.Genre, book.Stock)
		if err != nil {
			fmt.Printf("error %s inserting into the database\n", err)
			return false, err
		} else {
			fmt.Printf("successfully inserted %s into the database\n", book.Title)
			return true, nil
		}
	} else {
		fmt.Println("Book already exists")
		var ErrBookExists = errors.New("Book already exists")
		return false, ErrBookExists
	}
}

func UpdateBook(book types.BookInfo) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()

	// Check if the book exists before attempting to update
	isBookExists := checkBookExists(book)

	if isBookExists {
		updateSql := "UPDATE book_info SET Title = ?, Author = ?, Genre = ?, Stock = ? WHERE BookID = ?"
		_, err = db.Exec(updateSql, book.Title, book.Author, book.Genre, book.Stock, book.ID)
		if err != nil {
			fmt.Printf("error %s updating the database\n", err)
			return false, err
		} else {
			fmt.Printf("successfully updated book with ID %s in the database\n", strconv.Itoa(book.ID))
			return true, nil
		}
	} else {
		fmt.Println("Book does not exist")
		return false, nil
	}
}

func DeleteBook(book types.BookInfo) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()

	// Check if the book exists before attempting to delete
	isBookExists := checkBookExists(book)

	if isBookExists {
		deleteSql := "DELETE FROM book_info WHERE BookID=?"
		_, err = db.Exec(deleteSql, book.ID)
		if err != nil {
			fmt.Printf("error %s deleting from the database\n", err)
			return false, err
		} else {
			fmt.Printf("successfully deleted %s from the database\n", book.Title)
			return true, nil
		}
	} else {
		fmt.Println("Book does not exist")
		return false, nil
	}
}
