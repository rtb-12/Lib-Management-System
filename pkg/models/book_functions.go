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
func FetchBooks() types.ListBooks {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM book_info"
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchBooks []types.BookInfo
	for rows.Next() {
		var book types.BookInfo
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Stock)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchBooks = append(fetchBooks, book)
	}

	var listBooks types.ListBooks
	listBooks.Books = fetchBooks
	return listBooks

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

	insertSql := "INSERT INTO IssueRequest (UserID, BookID, Status) VALUES (?, ?, ?)"
	_, err = db.Exec(insertSql, info.UserID, info.BookID, types.Pending)
	if err != nil {
		fmt.Printf("error %s inserting into the database\n", err)
		return false, err
	} else {
		fmt.Printf("successfully inserted book issue request into the database\n")
		return true, nil
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
	_, err = db.Exec(inserSql, info.UserID, info.BookID, info.AdminID, info.IssueDate, info.ReturnDate)
	if err != nil {
		fmt.Printf("error %s inserting into the database\n", err)
		return false, err
	} else {
		updateSql := "UPDATE book_info SET Stock = Stock - 1 WHERE BookID = ?"
		_, err = db.Exec(updateSql, info.BookID)
		if err != nil {
			fmt.Printf("error %s updating the stock\n", err)
			return false, err
		}
		updateStatusSql := "UPDATE IssueRequest SET Status = ? WHERE UserID = ? AND BookID = ?"
		_, err = db.Exec(updateStatusSql, types.Approved, info.UserID, info.BookID)
		if err != nil {
			fmt.Printf("error %s updating the status\n", err)
			return false, err
		}
		fmt.Printf("successfully inserted book issue request into the database\n")
		return true, nil
	}

}

func ReturnBook(info types.BookReturn) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()
	inserSql := "UPDATE IssuedBook SET isReturned = ?, ReturnDate = ? WHERE IssuedId= ?"
	_, err = db.Exec(inserSql, true, info.ReturnDate, info.IssuedId)
	if err != nil {
		fmt.Printf("error %s updating the database for book return\n", err)
		return false, err
	} else {
		updateSql := "UPDATE book_info SET Stock = Stock + 1 WHERE BookID = (SELECT BookID FROM IssuedBook WHERE IssuedId = ?)"
		_, err = db.Exec(updateSql, info.IssuedId)
		if err != nil {
			fmt.Printf("error %s updating the stock\n", err)
			return false, err
		}
		fmt.Printf("successfully inserted book issue request into the database\n")
		return true, nil
	}

}

func RejectBookIssueRequest(info types.RejectBookIssueRequest) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database\n", err)
		return false, err
	}
	defer db.Close()
	inserSql := "UPDATE IssueRequest SET Status = ? WHERE UserID = ? AND BookID = ?"
	_, err = db.Exec(inserSql, types.Rejected, info.UserID, info.BookID)
	if err != nil {
		fmt.Printf("error %s updating the database\n", err)
		return false, err
	} else {
		fmt.Printf("successfully rejected book issue request into the database\n")
		return true, nil
	}

}

func FetchBookIssueRequests() types.ListBookIssueResponse {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM IssueRequest WHERE Status = 'pending'"
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var bookIssueRequests []types.BookIssueResponse
	for rows.Next() {
		var request types.BookIssueResponse
		err := rows.Scan(&request.RequestId, &request.UserID, &request.BookId, &request.Status)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		bookIssueRequests = append(bookIssueRequests, request)
	}

	var listBookIssueRequests types.ListBookIssueResponse
	listBookIssueRequests.BookIssueRequests = bookIssueRequests
	return listBookIssueRequests
}

func FetchBookIssuedBookIssued() types.ListBookIssued {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM IssuedBook WHERE isReturned = false"
	rows, err := db.Query(selectSql)
	db.Close()

	var bookIssuedBookList []types.BookIssuedDB
	for rows.Next() {
		var issued types.BookIssuedDB
		err := rows.Scan(&issued.IssuedId, &issued.UserID, &issued.BookID, &issued.AdminID, &issued.IssueDate, &issued.ReturnDate, &issued.IsReturned)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		bookIssuedBookList = append(bookIssuedBookList, issued)
	}

	var listBookIssued types.ListBookIssued
	listBookIssued.BooksIssued = bookIssuedBookList
	return listBookIssued

}

func FetchBookIssuedOfUser(userID types.UserID) types.ListBookIssued {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM IssuedBook WHERE UserID = ?"
	rows, err := db.Query(selectSql, userID.UserId)
	db.Close()

	var bookIssuedBookList []types.BookIssuedDB
	for rows.Next() {
		var issued types.BookIssuedDB
		err := rows.Scan(&issued.IssuedId, &issued.UserID, &issued.BookID, &issued.AdminID, &issued.IssueDate, &issued.ReturnDate, &issued.IsReturned)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		bookIssuedBookList = append(bookIssuedBookList, issued)
	}

	var listBookIssued types.ListBookIssued
	listBookIssued.BooksIssued = bookIssuedBookList
	return listBookIssued

}
