package main

import (
	"database/sql"
	"log"
	"strings"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func signUp(user User) error {
	// SQLite is a file based database.
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		log.Println(err)
		return err
	}

	defer sqliteDatabase.Close() // Defer Closing the database
	// createTableUser()                                                // Create Database Tables

	// INSERT RECORDS
	err = dbInsertUser(user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	// DISPLAY INSERTED RECORDS
	err = displayUser()
	if err != nil {
		return err
	}
	return nil
}

func loginSQL(user User) (bool, error) {
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Ouvre la base de données
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer sqliteDatabase.Close() // fermer la database quand on a fini
	ppl := User{}
	row, err := sqliteDatabase.Query("SELECT * FROM Customer WHERE UserName = ?", user.Username, user.Password) // Requete SQLite
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&ppl.Username, &ppl.Email, &ppl.Password, &ppl.UUID) // assigne chaque collone de la case a la structure de type User
	}
	if ppl.Username == user.Username && bcrypt.CompareHashAndPassword([]byte(ppl.Password), []byte(user.Password)) == nil { //compare le mot de passe avec le mot de passe crypté
		log.Println(ppl)
		log.Println("vous êtes connecté avec succès")
		return true, nil // si le mot de passe est bon
	} else {
		log.Println("Ce compte n'existe pas")
		return false, nil // si le mot de passe est faux
	}

}

func displayUser() error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Ouvre la base de données
	if err != nil {
		log.Println(err)
		return err
	}

	defer db.Close() // fermer la database quand on a fini

	row, err := db.Query("SELECT * FROM Customer ORDER BY UserName") //Cherche dans la base tout les utilisateurs par ordre alphabétique de nom
	if err != nil {
		log.Println(err)
		return err
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var UserName string
		var Email string
		var password string
		var UUID string
		row.Scan(&UserName, &Email, &password, &UUID) // assigne chaque collone de la table au variables
		log.Println("User: ", UserName, " ", Email, " ", password, " ", UUID)
	}
	return nil
}

func searchUser(UserName string) (User, error) {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Ouvre la base de données
	if err != nil {
		log.Println(err)
		return User{}, err
	}
	defer db.Close() // fermer la database quand on a fini

	row, err := db.Query("SELECT * FROM Customer WHERE UserName = ?", UserName) // Cherche dans la base l'utilisateur avec le nom d'utilisateur
	if err != nil {
		log.Println(err.Error() + " YA UNE ERREUR LA DANS SEARCH USER")
		return User{}, err
	}
	var ppl = User{}
	defer row.Close()
	for row.Next() {
		row.Scan(&ppl.Username, &ppl.Email, &ppl.Password, &ppl.UUID) // assigne chaque collone de la case a la structure de type User
	}

	return ppl, nil
}

func searchLikePost(id int) (string, error) {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer db.Close() // Fermer la database quand on a fini

	row, err := db.Query("SELECT * FROM TablePost WHERE ID = ?", id) // Cherche dans la base le post avec l'ID
	if err != nil {
		log.Println(err.Error() + " YA UNE ERREUR LA DANS SEARCH TALBE LIKE")
		return "", err
	}
	var ppl = Post{}
	defer row.Close()
	var strLike string
	for row.Next() {
		row.Scan(&ppl.ID, &ppl.Name, &ppl.Contentpost, &ppl.Categorie, &strLike, &ppl.UUID) // assigne chaque collone de la case a la structure de type Post
	}
	// IDLike := strings.Split(ppl.Like, " ")
	return strLike, nil
}

func postDB() ([]Post, error) {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	TablePost, err := db.Query("SELECT * FROM TablePost") // Cherche dans la base tout les posts
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var post Post
	var data []Post
	var strLike string
	for TablePost.Next() {
		TablePost.Scan(&post.ID, &post.Name, &post.Contentpost, &post.Categorie, &strLike, &post.UUID) // assigne chaque collone de la case a la structure de type Post
		post.Like = len(strings.Split(strLike, " ")) - 1                                               // -1 pour ne pas compter le vide
		data = append(data, post)                                                                      // ajoute le post a la liste
	}
	return data, nil
}

func triPost(Categorie string) ([]Post, error) {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM TablePost WHERE categorie = ?", Categorie) // Cherche dans la base tout les post de la catégorie demandé
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var post Post
	var data []Post
	defer row.Close()
	for row.Next() {
		row.Scan(&post.ID, &post.Name, &post.Contentpost, &post.Categorie, &post.Like, &post.UUID) // assigne chaque collone de la case a la structure de type Post
		data = append(data, post)
	}
	log.Println(data)
	return data, nil
}
