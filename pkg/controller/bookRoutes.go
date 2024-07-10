package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rtb-12/Lib-Management-System/pkg/models"
	"github.com/rtb-12/Lib-Management-System/pkg/types"
)

func Add(writer http.ResponseWriter, request *http.Request) {
	var book types.BookInfo
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookAdded, err := models.AddBook(book)
	if err != nil {
		fmt.Printf("Error adding book: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Error Adding the Book probably already exist", "BookName": "%s", "BookID": "%d", "Author":"%s"  }`, book.Title, book.ID, book.Author)
		writer.Write([]byte(response))
		return
	}
	if isBookAdded {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Added", "BookName": "%s", "BookID": "%d", "Author":"%s"  }`, book.Title, book.ID, book.Author)
		writer.Write([]byte(response))
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var book types.BookInfo
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookUpdated, err := models.UpdateBook(book)
	if err != nil {
		fmt.Printf("Error adding book: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBookUpdated {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Updated", "BookName": "%s", "BookID": "%d", "Author":"%s" , "Genre":"%s" , "Stock": "%d"  }`, book.Title, book.ID, book.Author, book.Genre, book.Stock)
		writer.Write([]byte(response))
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book types.BookInfo
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookDeleted, err := models.DeleteBook(book)
	if err != nil {
		fmt.Printf("Error deleting book: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBookDeleted {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Deleted", "BookName": "%s", "BookID": "%d", "Author":"%s" , "Genre":"%s" , "Stock": "%d"  }`, book.Title, book.ID, book.Author, book.Genre, book.Stock)
		w.Write([]byte(response))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func List(writer http.ResponseWriter, request *http.Request) {
	booksList := models.FetchBooks() // Assuming FetchBooks returns a slice of books

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(booksList); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RequestBookIssue(writer http.ResponseWriter, request *http.Request) {
	var issueRequest types.BookIssueRequest
	err := json.NewDecoder(request.Body).Decode(&issueRequest)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookRequested, err := models.CreateBookIssueRequest(issueRequest)
	if err != nil {
		fmt.Printf("Error requesting book: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBookRequested {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Requested", "BookName": "%s", "BookID": "%d", "UserId":"%d"  }`, issueRequest.Book.Title, issueRequest.Book.ID, issueRequest.UserID)
		writer.Write([]byte(response))
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}

}

func RejectBookIssueRequest(writer http.ResponseWriter, request *http.Request) {
	var issueRequest types.RejectBookIssueRequest
	err := json.NewDecoder(request.Body).Decode(&issueRequest)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookRequestRejected, err := models.RejectBookIssueRequest(issueRequest)
	if err != nil {
		fmt.Printf("Error rejecting book issue request: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBookRequestRejected {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Request Rejected",  "BookID": "%d", "UserId":"%d"   }`, issueRequest.BookID, issueRequest.UserID)
		writer.Write([]byte(response))
	} else {
		writer.WriteHeader(http.StatusBadRequest)

	}
}

func IssueBook(writer http.ResponseWriter, request *http.Request) {
	var issueRequest types.BookIssue
	err := json.NewDecoder(request.Body).Decode(&issueRequest)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookIssued, err := models.IssueBook(issueRequest)
	if err != nil {
		fmt.Printf("Error issuing book: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBookIssued {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Issued", "BookID": "%d", "UserId":"%d"   }`, issueRequest.BookID, issueRequest.UserID)
		writer.Write([]byte(response))
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func ReturnBookRequest(writer http.ResponseWriter, request *http.Request) {

	var issueRequest types.BookReturn
	err := json.NewDecoder(request.Body).Decode(&issueRequest)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	isBookReturned, err := models.ReturnBook(issueRequest)
	if err != nil {
		fmt.Printf("Error returning book: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBookReturned {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Book Returned", "IssuedID": "%d", "ReturnDate":"%s"   }`, issueRequest.IssuedId, issueRequest.ReturnDate)
		writer.Write([]byte(response))
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func FetchBookIssueRequests(writer http.ResponseWriter, request *http.Request) {
	booksIssueRequestsList := models.FetchBookIssueRequests()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(booksIssueRequestsList); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}

func FetchBookIssuedList(writer http.ResponseWriter, request *http.Request) {
	booksIssuedList := models.FetchBookIssuedBookIssued()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(booksIssuedList); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}
