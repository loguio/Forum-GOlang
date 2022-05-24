package main

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Username string
	Password string
	Email    string
}

type Post struct {
	Name string
	Contentpost string
	Categorie string
}

func signUp(user User) {
	// SQLite is a file based database.
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	// createTable(sqliteDatabase)                                      // Create Database Tables

	// INSERT RECORDS
	insertUser(sqliteDatabase, user.Username, user.Email, user.Password)

	// DISPLAY INSERTED RECORDS
	displayUser(sqliteDatabase)
}

func loginSQL(user User) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	// SELECT RECORDS
	ppl := User{}
	row := sqliteDatabase.QueryRow("SELECT * FROM Customer WHERE UserName = ? AND password = ?", user.Username, user.Password)
	// WHERE UserName = ? AND password = ?

	row.Scan(&ppl.Username, &ppl.Email, &ppl.Password)
	if ppl.Username == user.Username && ppl.Password == user.Password {
		fmt.Println(ppl)
		fmt.Println("vous êtes connecté avec succès")
	} else {
		fmt.Println("Ce compte n'existe pas")
	}

}

// func createTable(db *sql.DB) {
// 	createUserTableSQL := `CREATE TABLE IF NOT EXISTS Customer (
// 		"UserName" TEXT NOT NULL PRIMARY KEY,
// 		"Email" TEXT,
// 		"password" TEXT
// 	  );` // SQL Statement for Create Table

// 	log.Println("Create Customer table...")
// 	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	statement.Exec() // Execute SQL Statements
// 	log.Println("Users table created")
// }

// We are passing db reference connection from main to our method with other parameters
func insertUser(db *sql.DB, UserName string, Email string, password string) {
	log.Println("Inserting Users record ...")
	insertUserSQL := `INSERT or IGNORE INTO Customer(UserName, Email, password) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(UserName, Email, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayUser(db *sql.DB) {
	row, err := db.Query("SELECT * FROM Customer ORDER BY UserName")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var UserName string
		var Email string
		var password string
		row.Scan(&UserName, &Email, &password)
		log.Println("User: ", UserName, " ", Email, " ", password)
	}
}

func addpost(db *sql.DB, Name string, Contentpost string, Categorie string) {
	log.Println("Inserting new post ...")
	insertPostSQL := `INSERT or IGNORE INTO TablePost(name, contentpost, categorie) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertPostSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(Name ,Contentpost, Categorie)
	if err != nil {
		log.Fatalln(err.Error())
	}
	print("Post added")
}

// func createTable(db *sql.DB) {
// 	createPostTableSQL := `CREATE TABLE IF NOT EXISTS TablePost (
// 		"name" TEXT NOT NULL PRIMARY KEY,
// 		"contentpost" TEXT,
// 		"categorie" TEXT
// 	  );` // SQL Statement for Create Table

// 	log.Println("Create TablePost table...")
// 	statement, err := db.Prepare(createPostTableSQL) // Prepare SQL Statement
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	statement.Exec() // Execute SQL Statements
// 	log.Println("Post table created")
// }

func PostAdd(onePost Post) {
	// SQLite is a file based database.
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	// createTable(sqliteDatabase)                                      // Create Database Tables

	// INSERT RECORDS
	addpost(sqliteDatabase, onePost.Name,onePost.Contentpost, onePost.Categorie  )
}

func addbase(db *sql.DB) {
	insertPostSQL := `INSERT or IGNORE INTO TablePost(name, contentpost, categorie) VALUES (test, test, test)`
	statement, err := db.Prepare(insertPostSQL)
	if err != nil {
		print(statement)
	}
}