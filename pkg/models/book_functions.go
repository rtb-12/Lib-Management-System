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
		var errBookExists = errors.New("book already exists")
		return false, errBookExists
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

func CreateBookIssueRequest(info types.BookIssueRequest) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()

	// Check if the book exists before attempting to delete
	isBookExists := checkBookExists(info.Book)

	if isBookExists {
		insertSql := "INSERT INTO IssueRequest (UserID, BookID, Status) VALUES (?, ?, ?)"
		_, err = db.Exec(insertSql, info.UserID, info.Book.ID, info.Status)
		if err != nil {
			fmt.Printf("error %s inserting into the database\n", err)
			return false, err
		} else {
			fmt.Printf("successfully inserted book issue request into the database\n")
			return true, nil
		}
	} else {
		fmt.Println("Book does not exist")
		return false, nil
	}
}

func IssueBook(info types.BookIssue) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()
	inserSql := "INSERT INTO IssuedBook (UserID, BookID, AdminID, IssueDate,ReturnDate,isReturned) VALUES (?, ?, ?, ?,?,false)"
	_, err = db.Exec(inserSql, info.UserID, info.Book.ID, info.AdminID, info.IssueDate, info.ReturnDate)
	if err != nil {
		fmt.Printf("error %s inserting into the database\n", err)
		return false, err
	} else {
		updateSql := "UPDATE book_info SET Stock = Stock - 1 WHERE BookID = ?"
		_, err = db.Exec(updateSql, info.Book.ID)
		if err != nil {
			fmt.Printf("error %s updating the stock\n", err)
			return false, err
		}
		updateStatusSql := "UPDATE IssueRequest SET Status = ? WHERE UserID = ? AND BookID = ?"
		_, err = db.Exec(updateStatusSql, types.Approved, info.UserID, info.Book.ID)
		if err != nil {
			fmt.Printf("error %s updating the status\n", err)
			return false, err
		}
		fmt.Printf("successfully inserted book issue request into the database\n")
		return true, nil
	}

}

func ReturnBook(info types.BookIssue) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()
	inserSql := "UPDATE IssuedBook SET isReturned = ? WHERE UserID = ? AND BookID = ?"
	_, err = db.Exec(inserSql, true, info.UserID, info.Book.ID)
	if err != nil {
		fmt.Printf("error %s updating the database for book return\n", err)
		return false, err
	} else {
		updateSql := "UPDATE book_info SET Stock = Stock + 1 WHERE BookID = ?"
		_, err = db.Exec(updateSql, info.Book.ID)
		if err != nil {
			fmt.Printf("error %s updating the stock\n", err)
			return false, err
		}
		fmt.Printf("successfully inserted book issue request into the database\n")
		return true, nil
	}

}

func RejectBookIssueRequest(info types.BookIssueRequest) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()
	inserSql := "UPDATE book_info SET Status = ? WHERE UserID = ? AND BookID = ?"
	_, err = db.Exec(inserSql, types.Rejected, info.UserID, info.Book.ID)
	if err != nil {
		fmt.Printf("error %s updating the database\n", err)
		return false, err
	} else {
		fmt.Printf("successfully rejected book issue request into the database\n")
		return true, nil
	}

}
