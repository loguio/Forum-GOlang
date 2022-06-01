package main

import (
	"database/sql"
	"fmt"
	"log"
)

func createTableUser() {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

	// os.Create("./sqlite-database.db")
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS Customer (
		"UserName" TEXT NOT NULL,
		"Email" TEXT,
		"password" TEXT,
		"UUID" TEXT NOT NULL PRIMARY KEY
	  );` // SQL Statement for Create Table

	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func createTablePost() {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

	createPostTableSQL := `CREATE TABLE IF NOT EXISTS TablePost (
		"ID" INTEGER PRIMARY KEY AUTOINCREMENT,
		"Name" TEXT NOT NULL,
		"Contentpost" TEXT NOT NULL,
		"Categorie" TEXT,
		"Like" TEXT,
		"UUID_User" TEXT NOT NULL,
		FOREIGN KEY("UUID_User") REFERENCES Customer("UUID")
	  );` // SQL Statement for Create Table

	log.Println("Create TablePost table...")
	statement, err := db.Prepare(createPostTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error() + " ICI PB FOREIGN KEY")
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Post table created")
}
