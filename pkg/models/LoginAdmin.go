package models

import (
	"database/sql"
	"fmt"

	"github.com/rtb-12/Lib-Management-System/pkg/types"
	"github.com/rtb-12/Lib-Management-System/pkg/utils"
)

func CheckAdmin(admin types.AdminLogin) (int, error) {
	db, err := Connection()
	if err != nil {
		return 0, fmt.Errorf("error %s connecting to the database", err)
	}
	defer db.Close()

	// Prepare the SQL statement to select the admin's hashed password
	stmt, err := db.Prepare("SELECT AdminID, Password_Hash FROM Admin WHERE Email = ?")
	if err != nil {
		return 0, fmt.Errorf("error preparing SQL statement: %s", err)
	}

	var adminID int
	var storedHash string
	err = stmt.QueryRow(admin.Email).Scan(&adminID, &storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// Admin does not exist
			return 0, fmt.Errorf("Admin does not exist")
		}
		// An error occurred during query execution
		return 0, fmt.Errorf("error executing statement: %s", err)
	}

	// Compare the provided password with the stored hash
	match := utils.CompareHashPassword(admin.Password, storedHash)
	if !match {
		return 0, fmt.Errorf("invalid login credentials")
	}

	return adminID, nil
}
