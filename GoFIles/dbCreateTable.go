package main

import (
	"database/sql"
	"fmt"
	"log"
)

func createTableUser() { //fonction pour créer la table User en cas de besoin
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Ouverture de la database
	defer db.Close()                                     //ferme la database
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS Customer (
		"UserName" TEXT NOT NULL,
		"Email" TEXT,
		"password" TEXT,
		"UUID" TEXT NOT NULL PRIMARY KEY
	  );` // SQL Statement for Create Table

	statement, err := db.Prepare(createUserTableSQL) // Prepare la requete sqlite
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func createTablePost() { //fonction pour créer la table Post en cas de besoin
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Ouverture de la database
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
	statement, err := db.Prepare(createPostTableSQL) // prepare la requete sqlite
	if err != nil {
		log.Fatal(err.Error() + " ICI PB FOREIGN KEY")
	}
	statement.Exec() // Execute la requete sqlite
	log.Println("Post table created")
}
