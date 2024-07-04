package models

import (
	"database/sql"
	"fmt"

	"github.com/rtb-12/Lib-Management-System/pkg/types"
	"github.com/rtb-12/Lib-Management-System/pkg/utils"
)

func CheckUser(user types.UserLogin) (int, error) {
	db, err := Connection()
	if err != nil {
		return 0, fmt.Errorf("error %s connecting to the database", err)
	}
	defer db.Close()

	// Prepare the SQL statement to select the user's hashed password
	stmt, err := db.Prepare("SELECT UserID, Password_Hash FROM User WHERE Email = ?")
	if err != nil {
		return 0, fmt.Errorf("error preparing statement: %s", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	var userID int
	var storedHash string
	err = stmt.QueryRow(user.Email).Scan(&userID, &storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// User does not exist
			return 0, fmt.Errorf("User does not exist")
		}
		// An error occurred during query execution
		return 0, fmt.Errorf("error executing statement: %s", err)
	}

	// Compare the provided password with the stored hash
	match := utils.CompareHashPassword(user.Password, storedHash)
	if !match {
		return 0, fmt.Errorf("invalid login credentials")
	}

	// fmt.Println("User is valid, ID:", userID)
	return userID, nil
}
