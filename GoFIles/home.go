package main

import (
	"html/template"
	"log"
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
		erreur500(w)
		return
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

				log.Println("buttonSelect : " + buttonSelect)
				tri, err = triPost(buttonSelect) //tri les post par categorie
				if err != nil {
					erreur500(w)
					return
				}
				data.Posts = tri //affiche les post trié

			} else if Name != "" && Contentpost != "" && Categorie != "" { //si il y a un post

				post = Post{Name: Name, Contentpost: Contentpost, Categorie: Categorie}
				// createTablePost()
				err = dbInsertPost(post.Name, post.Contentpost, post.Categorie, cookie.Value) //ajoute le post dans la base
				if err != nil {
					erreur500(w)
					return
				}
				data.Posts, err = postDB() //recupere les post de la base pour les afficher
				if err != nil {
					erreur500(w)
					return
				}

			} else if IDButtonLike != "" { //si il y a un like

				log.Println("buttonLike : " + IDButtonLike)
				intButtonLike, err := strconv.Atoi(IDButtonLike) //converti le string en int
				if err != nil {
					log.Println("error : ", err)
					erreur500(w)
					return
				}
				err = dbLike(intButtonLike, cookie.Value) //ajoute le like dans la base
				if err != nil {
					erreur500(w)
					return
				}
				data.Posts, err = postDB() //recupere les post de la base pour les afficher
				if err != nil {
					erreur500(w)
					return
				}
			} else {
				// createTablePost()
				data.Posts, err = postDB() //recupere les post de la base pour les afficher
				if err != nil {
					erreur500(w)
					return
				}
			}
		}
	} else {
		data.Connected = false //non connecté
		buttonSelect := r.FormValue("buttonCategorie")
		if buttonSelect != "" && buttonSelect != "all" { //si il y a un bouton de categorie
			log.Println("buttonSelect : " + buttonSelect)
			tri, err = triPost(buttonSelect) //tri les post par categorie
			if err != nil {
				erreur500(w)
				return
			}
			data.Posts = tri //affiche les post trié

		} else {
			data.Posts, err = postDB() //recupere les post de la base pour les afficher
			if err != nil {
				erreur500(w)
				return
			}
		}
	}

	tmpl.ExecuteTemplate(w, "home", data) //execute le template home.html
}
