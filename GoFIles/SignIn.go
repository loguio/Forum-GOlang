package main

import (
	"html/template"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("../template/signIn.html") // utilisation du fichier signIn.html sur le template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == "POST" { // si la requête est de type POST
		UserName := r.FormValue("Username")         // récupère le username
		Password := r.FormValue("password")         // récupère le password
		Email := r.FormValue("Mail")                // récupère l'email
		confPassword := r.FormValue("confPassword") // récupère la confirmation du password
		if Password == confPassword {               // si les mots de passe sont identiques
			user = User{Username: UserName, Password: Password, Email: Email}
			signUp(user) // enregistre le nouvel utilisateur
		} //sinon ne fais rien
	}
	tmpl.ExecuteTemplate(w, "Signin", user) // execute le template
}
