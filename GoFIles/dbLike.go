package main

import (
	"database/sql"
	"log"
	"strings"
)

func remove(s []string, i int) []string { //fonction pour enlever un élément d'un tableau
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func dbLike(id int, UUID string) error {
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()                //ferme la database
	like, err := searchLikePost(id) //appel la fonction searchLikePost
	if err != nil {
		return err
	}
	var there bool                     //variable pour savoir si l'utilisateur a déjà liké le post
	IDLike := strings.Split(like, " ") //split le string pour récupérer les ID des posts likés
	var i int
	for i = 0; i < len(IDLike); i++ { //boucle pour vérifier si l'utilisateur a déjà liké le post
		if IDLike[i] == UUID {
			there = true
			break
		}
	}
	if there { //si l'utilisateur a déjà liké le post
		IDLike = remove(IDLike, i)                              //enlève le post de la liste
		like = strings.Join(IDLike, " ")                        //reconstruit le string
		addLike := `UPDATE TablePost SET Like = ? WHERE id = ?` //création de la requête sqlite
		statement, err := db.Prepare(addLike)                   //prepare la requête
		if err != nil {
			log.Println(err.Error())
			return err
		}
		_, err = statement.Exec(like, id) //execute la requête
		if err != nil {
			log.Println(err.Error())
			return err
		}
		log.Println("like removed")
	} else {
		addLike := `UPDATE TablePost SET Like = ? WHERE id = ?` //création de la requête sqlite
		statement, err := db.Prepare(addLike)                   //prepare la requête
		if err != nil {
			log.Println(err.Error())
			return err
		}
		_, err = statement.Exec(like+" "+UUID, id) //execute la requête pour ajouter un like
		if err != nil {
			log.Println(err.Error())
			return err
		}
		log.Println("like added")
	}
	return nil
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
