package main

import (
	"html/template"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("../template/signIn.html") // utilisation du fichier navPage.gohtml pour le template
	if r.Method == "POST" {
		UserName := r.FormValue("Username")
		Password := r.FormValue("password")
		Email := r.FormValue("Mail")
		confPassword := r.FormValue("confPassword")
		if Password == confPassword {
			user = User{Username: UserName, Password: Password, Email: Email}
			signUp(user)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// database()
	}
	tmpl.ExecuteTemplate(w, "Signin", user)
}
