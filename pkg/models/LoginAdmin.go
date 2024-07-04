package models

import (
	"database/sql"
	"fmt"

	// Add this line to import the package
	"github.com/rtb-12/Lib-Management-System/pkg/types"
	"github.com/rtb-12/Lib-Management-System/pkg/utils"
)

func CheckAdmin(admin types.AdminLogin) (int, error) {
	db, err := Connection()
	if err != nil {
		return 0, fmt.Errorf("error %s connecting to the database", err)
	}
	defer db.Close()

	//Hash the admin's password
	hashedPassword, err := utils.GenerateHashPassword(admin.Password)
	if err != nil {
		return 0, fmt.Errorf("error hashing password: %s", err)
	}

	//Prepare the SQL statement to select the admin ID
	stmt, err := db.Prepare("SELECT AdminID FROM Admin WHERE Email = ? AND Password_Hash = ?")
	if err != nil {
		return 0, fmt.Errorf("error preparing SQL statement: %s", err)
	}

	//Prepare the SQL statement
	var adminID int
	err = stmt.QueryRow(admin.Email, hashedPassword).Scan(&adminID)
	if err != nil {
		if err == sql.ErrNoRows {
			//Admin does not exist
			fmt.Errorf("Admin does not exist")
			return 0, nil
		}
		//An error occurred during query execution
		return 0, fmt.Errorf("error executing statement: %s", err)
	}

	fmt.Println("Admin is valid, ID:", adminID)
	return adminID, nil
}
