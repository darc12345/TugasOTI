package repository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (h *DBrepo) PostLoginRepo(givenUserId string, passwordhash string) (string, string, error) {
	var Id string
	var Role string
	db := h.db

	// Query for user ID and role based on credentials
	err := db.QueryRow("SELECT USERID,ROLE FROM Users WHERE USERID=? AND PASSWORDHASH=?;", givenUserId, passwordhash).Scan(&Id, &Role)
	if err != nil {
		log.Printf("Failed to retrieve user login information. User ID: %s. Original Error: %s", givenUserId, err)
		return "", "", fmt.Errorf("error during login attempt")
	}

	return Id, Role, nil
}

func (h *DBrepo) PostRegisterRepo(email string, address string, passwordhash string) error {
	db := h.db

	// Insert new user record into the database
	_, err := db.Exec("INSERT INTO Users(USERID, EMAIL, ADDRESS, PASSWORDHASH) VALUES(?,?,?,?)",
		uuid.New().String(), email, address, passwordhash)
	if err != nil {
		log.Printf("Failed to register new user. Email: %s. Original Error: %s", email, err)
		return fmt.Errorf("error during registration")
	}

	return nil
}
