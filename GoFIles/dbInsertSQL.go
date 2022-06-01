package main

import (
	"database/sql"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// We are passing db reference connection from main to our method with other parameters
func dbInsertUser(UserName string, Email string, password string) {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

	log.Println("Inserting Users record ...")
	insertUserSQL := `INSERT or IGNORE INTO Customer(UUID, UserName, Email, password) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	var er = error(nil)
	u1 := uuid.Must(uuid.NewV4(), er)
	passwordCrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err = statement.Exec(u1, UserName, Email, string(passwordCrypt)) // Execute SQL Statement
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("User added")
}

func dbInsertPost(Name string, Contentpost string, Categorie string, cookieValue string) {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

	log.Println("Inserting new post ...")
	insertPostSQL := `INSERT or IGNORE INTO TablePost(Name, Contentpost, Categorie, UUID_User) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertPostSQL)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(Name, Contentpost, Categorie, cookieValue)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Post added")
}
