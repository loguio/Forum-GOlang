package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	_ "github.com/satori/go.uuid"
)

func main() {
	fileServer := http.FileServer(http.Dir("static/")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", handler) // lance l'erreur 404 quand on est sur une URL pas utilisée

	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/Login":
		Login(w, r)
	case "/Signin":
		Signin(w, r)
	case "/home":
		home(w, r)
	case "/logout":
		Logout(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not implemented")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../template/Log.html")
	cookie, err := r.Cookie("session-id")
	if err == nil {
		fmt.Println("cookie value before logout : " + cookie.Value)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		fmt.Println("You've successfully logged out.")

	} else {
		fmt.Println("You're not logged in.")
	}

	tmpl.ExecuteTemplate(w, "Login", nil)

}

func profile(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./profile.html")
	if err != nil {
	}
	tmpl.ExecuteTemplate(w, "profile", user)
}
