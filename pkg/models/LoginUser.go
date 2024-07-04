package models

import (
	"database/sql"
	"fmt"

	// Add this line to import the package
	"github.com/rtb-12/Lib-Management-System/pkg/types"
	"github.com/rtb-12/Lib-Management-System/pkg/utils"
)

func CheckUser(user types.UserLogin) (int, error) {
	db, err := Connection()
	if err != nil {
		return 0, fmt.Errorf("error %s connecting to the database", err)
	}
	defer db.Close()

	// Hash the user's password
	hashedPassword, err := utils.GenerateHashPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("error hashing password: %s", err)
	}

	// Prepare the SQL statement to select the user ID
	stmt, err := db.Prepare("SELECT UserID FROM User WHERE Email = ? AND Password_Hash = ?")
	if err != nil {
		return 0, fmt.Errorf("error preparing statement: %s", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	var userID int
	err = stmt.QueryRow(user.Email, hashedPassword).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// User does not exist
			fmt.Errorf("User does not exist")
			return 0, nil
		}
		// An error occurred during query execution
		return 0, fmt.Errorf("error executing statement: %s", err)
	}

	fmt.Println("User is valid, ID:", userID)
	return userID, nil
}
