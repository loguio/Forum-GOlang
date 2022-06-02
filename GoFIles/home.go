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
		data.Connected = true      //connecté
		data.Posts, err = postDB() //recupere les post de la base pour les afficher
		log.Println("data.Posts : ", data.Posts)
		if err != nil {
			erreur500(w)
			return
		}
		if r.Method == "POST" {

			Name := r.FormValue("Name")
			Contentpost := r.FormValue("Contentpost")
			Categorie := r.FormValue("Categorie")
			buttonSelect := r.FormValue("buttonCategorie")
			IDButtonLike := r.FormValue("buttonLike")
			comment := r.FormValue("comment")
			idComment := r.FormValue("IDbuttonComment")
			IDButtonLikeCommnt := r.FormValue("buttonLikeComment")

			if buttonSelect != "" && buttonSelect != "all" { //si il y a un bouton de categorie

				log.Println("buttonSelect : " + buttonSelect)
				tri, err = triPost(buttonSelect) //tri les post par categorie
				if err != nil {
					erreur500(w)
					return
				}
				data.Posts = tri //affiche les post trié

			} else if Name != "" && Contentpost != "" && Categorie != "" { //si il y a un post
				log.Println("Tu as mis un post")
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
			} else if comment != "" && idComment != "" { //si il y a un commentaire
				log.Println("Tu as mis un commentaire")
				intIDComment, err := strconv.Atoi(idComment) //converti le string en int
				if err != nil {
					log.Println("error : ", err)
					erreur500(w)
					return
				}
				err = dbComment(comment, cookie.Value, intIDComment) //ajoute le commentaire dans la base
				if err != nil {
					erreur500(w)
					return
				}
				data.Posts, err = postDB() //recupere les post de la base pour les afficher
				if err != nil {
					erreur500(w)
					return
				}
			} else if IDButtonLikeCommnt != "" { //si il y a un like sur un commentaire
				log.Println("buttonLikeComment : " + IDButtonLikeCommnt)
				intButtonLikeCommnt, err := strconv.Atoi(IDButtonLikeCommnt) //converti le string en int
				if err != nil {
					log.Println("error : ", err)
					erreur500(w)
					return
				}
				err = dbLikeComment(intButtonLikeCommnt, cookie.Value) //ajoute le like dans la base
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
				log.Println("data.Posts : ", data.Posts)
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
