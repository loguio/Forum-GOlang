package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("../template/log.html") //charge le template
	if err != nil {
		erreur500(w)
	}
	if r.Method == "POST" { // si la requête est de type POST
		cookie, err := r.Cookie("session-id") //récupère le cookie session-id de la page
		if err == nil {                       // si il y a un cookie
			log.Println("You've already logged in.")
		} else if err != nil { //Si il y a pas de cookie
			log.Println("cookie not found")
			cookie = &http.Cookie{ //crée un nouveau cookie
				Name: "session-id",
			}

			UserName := r.FormValue("Username")   //récupère le username
			Password := r.FormValue("password")   //récupère le password
			if UserName != "" && Password != "" { //si les champs sont remplis
				user = User{Username: UserName, Password: Password}
				connected, err := loginSQL(user) //appel la fonction loginSQL
				if err != nil {
					erreur500(w)
					return
				}
				if connected == true { //si l'utilisateur est connecté
					sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
					defer sqliteDatabase.Close()                                     //ferme la database
					user, err = searchUser(UserName)                                 //appel la fonction searchUser
					if err != nil {
						erreur500(w)
						return
					}

					cookie.Value = user.UUID //applique le cookie avec le UUID de l'utilisateur
					cookie.MaxAge = 604800   //durée de validité du cookie

					http.SetCookie(w, cookie) //envoie le cookie sur la page

					w.WriteHeader(http.StatusOK) //envoie la réponse OK
				} else { //sinons

				}
			} else {
				log.Println("entrée vide")
			}
		}
	}
	log.Println("---------------------------------")
	tmpl.ExecuteTemplate(w, "Login", user) //execute le template
}
