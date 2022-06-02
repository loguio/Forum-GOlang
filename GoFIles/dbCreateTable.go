package main

import (
	"database/sql"
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
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func createTablePost() error { //fonction pour créer la table Post en cas de besoin
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Ouverture de la database
	if err != nil {
		log.Println(err)
		return err
	}
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
		log.Println(err.Error() + " ICI PB FOREIGN KEY")
		return err
	}
	statement.Exec() // Execute la requete sqlite
	log.Println("Post table created")
	return nil
}

func createTableComment() error { //fonction pour créer la table Post en cas de besoin
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Ouverture de la database
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	createPostTableSQL := `CREATE TABLE IF NOT EXISTS TableComment (
		"ID" INTEGER PRIMARY KEY AUTOINCREMENT,
		"Contentpost" TEXT NOT NULL,
		"Like" TEXT,
		"UUID_User" TEXT NOT NULL,
		"UUID_Post" TEXT NOT NULL,
		FOREIGN KEY("UUID_Post") REFERENCES TablePost("ID"),
		FOREIGN KEY("UUID_User") REFERENCES Customer("UUID")
	  );` // SQL Statement for Create Table

	log.Println("Create TableComment table...")
	statement, err := db.Prepare(createPostTableSQL) // prepare la requete sqlite
	if err != nil {
		log.Println(err.Error() + " ICI Table Post")
		return err
	}
	statement.Exec() // Execute la requete sqlite
	log.Println("Commentary table created")
	return nil
}
