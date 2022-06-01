package main

import (
	"database/sql"
	"log"

	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func dbInsertUser(UserName string, Email string, password string) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close() //ferme la database

	log.Println("Inserting Users record ...")
	insertUserSQL := `INSERT or IGNORE INTO Customer(UUID, UserName, Email, password) VALUES (?, ?, ?, ?)` //création de la requête sqlite
	statement, err := db.Prepare(insertUserSQL)                                                            // Prepare la requete
	// This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
		return err
	}
	var er = error(nil)
	u1 := uuid.Must(uuid.NewV4(), er)                                                       //création d'un UUID
	passwordCrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //cryptage du mot de passe
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = statement.Exec(u1, UserName, Email, string(passwordCrypt)) // Execute SQL Statement*
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("User added")
	return nil
}

func dbInsertPost(Name string, Contentpost string, Categorie string, cookieValue string) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close() //ferme la database

	log.Println("Inserting new post ...")
	insertPostSQL := `INSERT or IGNORE INTO TablePost(Name, Contentpost, Categorie, UUID_User) VALUES (?, ?, ?, ?)` //création de la requête sqlite
	statement, err := db.Prepare(insertPostSQL)                                                                     // Prepare la requete
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = statement.Exec(Name, Contentpost, Categorie, cookieValue) // Execute la requete sqlite
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Post added")
	return nil
}
