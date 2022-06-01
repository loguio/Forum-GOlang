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

	tmpl, err := template.ParseFiles("../template/home.html")
	if err != nil {
	}

	cookie, err := r.Cookie("session-id")
	if err == nil {
		data.Connected = true
		if r.Method == "POST" {

			Name := r.FormValue("Name")
			Contentpost := r.FormValue("Contentpost")
			Categorie := r.FormValue("Categorie")
			buttonSelect := r.FormValue("buttonCategorie")
			IDButtonLike := r.FormValue("buttonLike")

			if buttonSelect != "" && buttonSelect != "all" {

				fmt.Println("buttonSelect : " + buttonSelect)
				tri = triPost(buttonSelect)
				data.Posts = tri

			} else if Name != "" && Contentpost != "" && Categorie != "" {

				post = Post{Name: Name, Contentpost: Contentpost, Categorie: Categorie}
				createTablePost()
				dbInsertPost(post.Name, post.Contentpost, post.Categorie, cookie.Value)
				data.Posts = postDB()

			} else if IDButtonLike != "" {

				fmt.Println("buttonLike : " + IDButtonLike)
				intButtonLike, err := strconv.Atoi(IDButtonLike)
				if err != nil {
					fmt.Println("error : ", err)
				}
				dbLike(intButtonLike, cookie.Value)
				data.Posts = postDB()
			} else {
				// createTablePost()
				data.Posts = postDB()
				fmt.Println(data.Posts)
			}
		}
	} else {
		data.Connected = false
		buttonSelect := r.FormValue("buttonCategorie")
		if buttonSelect != "" && buttonSelect != "all" {
			fmt.Println("buttonSelect : " + buttonSelect)
			tri = triPost(buttonSelect)
			data.Posts = tri

		} else {
			data.Posts = postDB()
		}
	}

	tmpl.ExecuteTemplate(w, "home", data)
}
