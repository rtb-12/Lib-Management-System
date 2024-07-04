package models

import (
	"fmt"

	// Add this line to import the package
	"github.com/rtb-12/Lib-Management-System/pkg/types"
	"github.com/rtb-12/Lib-Management-System/pkg/utils"
)

func RegisterAdmin(admin types.Admin) error {
	db, err := Connection()
	if err != nil {
		return fmt.Errorf("error %s connecting to the database", err)
	}
	defer db.Close()

	// Hash the user's password
	hashedPassword, err := utils.GenerateHashPassword(admin.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %s", err)
	}

	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO Admin(Name, Email, Password_Hash) VALUES(?, ?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %s", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(admin.Name, admin.Email, hashedPassword)
	if err != nil {
		return fmt.Errorf("error executing statement: %s", err)
	}

	fmt.Println("Admin registered successfully")
	return nil
}
