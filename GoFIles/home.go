package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	var post Post
	var data data
	var tri []Post
	tri = nil

	tmpl, err := template.ParseFiles("../template/home.html") //parse le template home.html
	if err != nil {
	}

	cookie, err := r.Cookie("session-id") //recupere le cookie session-id
	if err == nil {                       //si il y a un cookie
		data.Connected = true //connecté
		if r.Method == "POST" {

			Name := r.FormValue("Name")
			Contentpost := r.FormValue("Contentpost")
			Categorie := r.FormValue("Categorie")
			buttonSelect := r.FormValue("buttonCategorie")
			IDButtonLike := r.FormValue("buttonLike")

			if buttonSelect != "" && buttonSelect != "all" { //si il y a un bouton de categorie

				fmt.Println("buttonSelect : " + buttonSelect)
				tri = triPost(buttonSelect) //tri les post par categorie
				data.Posts = tri            //affiche les post trié

			} else if Name != "" && Contentpost != "" && Categorie != "" { //si il y a un post

				post = Post{Name: Name, Contentpost: Contentpost, Categorie: Categorie}
				// createTablePost()
				dbInsertPost(post.Name, post.Contentpost, post.Categorie, cookie.Value) //ajoute le post dans la base
				data.Posts = postDB()                                                   //recupere les post de la base pour les afficher

			} else if IDButtonLike != "" { //si il y a un like

				fmt.Println("buttonLike : " + IDButtonLike)
				intButtonLike, err := strconv.Atoi(IDButtonLike) //converti le string en int
				if err != nil {
					fmt.Println("error : ", err)
				}
				dbLike(intButtonLike, cookie.Value) //ajoute le like dans la base
				data.Posts = postDB()               //recupere les post de la base pour les afficher
			} else {
				// createTablePost()
				data.Posts = postDB() //recupere les post de la base pour les afficher
				fmt.Println(data.Posts)
			}
		}
	} else {
		data.Connected = false //non connecté
		buttonSelect := r.FormValue("buttonCategorie")
		if buttonSelect != "" && buttonSelect != "all" { //si il y a un bouton de categorie
			fmt.Println("buttonSelect : " + buttonSelect)
			tri = triPost(buttonSelect) //tri les post par categorie
			data.Posts = tri            //affiche les post trié

		} else {
			data.Posts = postDB() //recupere les post de la base pour les afficher
		}
	}

	tmpl.ExecuteTemplate(w, "home", data) //execute le template home.html
}
