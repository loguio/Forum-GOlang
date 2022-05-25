package main

import (
	"database/sql"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
	Email    string
	UUID     string
}

type Post struct {
	Name        string
	Contentpost string
	Categorie   string
}

func signUp(user User) {
	// SQLite is a file based database.
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	createTable(sqliteDatabase)                                      // Create Database Tables

	// INSERT RECORDS
	insertUser(sqliteDatabase, user.Username, user.Email, user.Password)

	// DISPLAY INSERTED RECORDS
	displayUser(sqliteDatabase)
}

func loginSQL(user User) bool {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	// SELECT RECORDS
	ppl := User{}
	row, err := sqliteDatabase.Query("SELECT * FROM Customer WHERE UserName = ?", user.Username, user.Password)
	if err != nil {
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&ppl.Username, &ppl.Email, &ppl.Password, &ppl.UUID)
	}
	if ppl.Username == user.Username && bcrypt.CompareHashAndPassword([]byte(ppl.Password), []byte(user.Password)) == nil {
		fmt.Println(ppl)
		fmt.Println("vous êtes connecté avec succès")
		return true
	} else {
		fmt.Println("Ce compte n'existe pas")
		return false
	}

}

func createTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS Customer (
		"UserName" TEXT NOT NULL PRIMARY KEY,
		"Email" TEXT,
		"password" TEXT,
		"UUID" TEXT
	  );` // SQL Statement for Create Table

	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

// We are passing db reference connection from main to our method with other parameters
func insertUser(db *sql.DB, UserName string, Email string, password string) {
	log.Println("Inserting Users record ...")
	insertUserSQL := `INSERT or IGNORE INTO Customer(UserName, Email, password, UUID) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	var er = error(nil)
	u1 := uuid.Must(uuid.NewV4(), er)
	passwordCrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err = statement.Exec(UserName, Email, string(passwordCrypt), u1) // Execute SQL Statement
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
		var UUID string
		row.Scan(&UserName, &Email, &password, &UUID)
		log.Println("User: ", UserName, " ", Email, " ", password, " ", UUID)
	}
}

func searchUser(db *sql.DB, UserName string) User {
	row, err := db.Query("SELECT * FROM Customer WHERE UserName = ?", UserName)
	if err != nil {
	}
	var ppl = User{}
	defer row.Close()
	for row.Next() {
		row.Scan(&ppl.Username, &ppl.Email, &ppl.Password, &ppl.UUID)
	}

	return ppl
}
func addpost(db *sql.DB, Name string, Contentpost string, Categorie string) {
	log.Println("Inserting new post ...")
	insertPostSQL := `INSERT or IGNORE INTO TablePost(name, contentpost, categorie) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertPostSQL)
	if err != nil {
		// log.Fatalln(err.Error())
	}
	_, err = statement.Exec(Name, Contentpost, Categorie)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// func createTable(db *sql.DB) {
// 	createPostTableSQL := `CREATE TABLE IF NOT EXISTS TablePost (
// 		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
// 		"name" TEXT NOT NULL,
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

func postDB() []Post {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	TablePost, err := db.Query("SELECT * FROM TablePost")
	if err != nil {
		fmt.Println(err)
	}
	var post Post
	var data []Post
	for TablePost.Next() {
		TablePost.Scan(&post.Name, &post.Contentpost, &post.Categorie)
		data = append(data, post)
	}
	return data
}
