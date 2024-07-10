package database

import (
	"database/sql"
	"log"
)

func CheckEmailUnique(db *sql.DB, email string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		log.Println("Error checking email uniqueness: ", err)
		return false, err
	}
	return count == 0, nil
}
