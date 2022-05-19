package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type User struct {
	Username string
	Password string
	Email    string
}

func main() {
	fileServer := http.FileServer(http.Dir("static/")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/Login", Login)   // lance l'erreur 404 quand on est sur une URL pas utilisée
	http.HandleFunc("/Signin", Signin) // lance l'erreur 404 quand on est sur une URL pas utilisée
	http.HandleFunc("/home", home)
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}

func Signin(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./sign in.html") // utilisation du fichier navPage.gohtml pour le template
	if r.Method == "POST" {
		UserName := r.FormValue("Username")
		Password := r.FormValue("password")
		Email := r.FormValue("Mail")
		confPassword := r.FormValue("confPassword")
		if Password == confPassword {
			user = User{Username: UserName, Password: Password, Email: Email}
			database(user)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// database()
	}
	tmpl.ExecuteTemplate(w, "Signin", user)
}

func database(user User) {
	// SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("Yforum.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	createTable(sqliteDatabase)                                      // Create Database Tables

	// INSERT RECORDS
	insertUser(sqliteDatabase, user.Username, user.Email, user.Password)

	// DISPLAY INSERTED RECORDS
	displayUser(sqliteDatabase)
}
func createTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS Customer (
		"UserName" TEXT NOT NULL PRIMARY KEY,		
		"Email" TEXT,
		"password" TEXT	
	  );` // SQL Statement for Create Table

	log.Println("Create Customer table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Users table created")
}

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
func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./Log.html")
	if err != nil {
	}
	tmpl.ExecuteTemplate(w, "Login", user)
}

func home(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./home.html")
	if err != nil {
	}
	tmpl.ExecuteTemplate(w, "home", user)
}
