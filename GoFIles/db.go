package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func signUp(user User) {
	// SQLite is a file based database.
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	// createTableUser()                                                // Create Database Tables

	// INSERT RECORDS
	dbInsertUser(user.Username, user.Email, user.Password)

	// DISPLAY INSERTED RECORDS
	displayUser()
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

func displayUser() {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

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

func searchUser(UserName string) User {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

	row, err := db.Query("SELECT * FROM Customer WHERE UserName = ?", UserName)
	if err != nil {
		fmt.Println(err.Error() + " YA UNE ERREUR LA DANS SEARCH USER")
	}
	var ppl = User{}
	defer row.Close()
	for row.Next() {
		row.Scan(&ppl.Username, &ppl.Email, &ppl.Password, &ppl.UUID)
	}

	return ppl
}

func searchLikePost(id int) string {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

	row, err := db.Query("SELECT * FROM TablePost WHERE ID = ?", id)
	if err != nil {
		fmt.Println(err.Error() + " YA UNE ERREUR LA DANS SEARCH TALBE LIKE")
	}
	var ppl = Post{}
	defer row.Close()
	var strLike string
	for row.Next() {
		row.Scan(&ppl.ID, &ppl.Name, &ppl.Contentpost, &ppl.Categorie, &strLike, &ppl.UUID)
	}
	// IDLike := strings.Split(ppl.Like, " ")
	return strLike
}

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
	var strLike string
	for TablePost.Next() {
		TablePost.Scan(&post.ID, &post.Name, &post.Contentpost, &post.Categorie, &strLike, &post.UUID)
		post.Like = len(strings.Split(strLike, " ")) - 1
		data = append(data, post)
	}
	return data
}

func triPost(Categorie string) []Post {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

	row, err := db.Query("SELECT * FROM TablePost WHERE categorie = ?", Categorie)
	if err != nil {
	}
	var post Post
	var data []Post
	defer row.Close()
	for row.Next() {
		row.Scan(&post.ID, &post.Name, &post.Contentpost, &post.Categorie, &post.Like, &post.UUID)
		data = append(data, post)
	}
	fmt.Println(data)
	return data
}
