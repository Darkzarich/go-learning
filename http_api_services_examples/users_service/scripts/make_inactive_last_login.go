package main

import (
	"database/sql"
	"log"

	config "users-service/configs"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.Load()

	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	defer db.Close()

	log.Println("Updating last login for user 1")

	result, err := db.Exec("UPDATE users SET last_login = datetime('now', '-30 days') WHERE id = 1")
	if err != nil {
		log.Fatalf("did not update last login: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("could not get rows affected: %v", err)
	}

	log.Printf("Updated %d rows\n", rowsAffected)
}
