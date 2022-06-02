package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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
	case "/profile":
		profile(w, r)
	case "/logout":
		Logout(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not implemented")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../template/log.html")
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
	cookie, err := r.Cookie("session-id")
	data := DataProfile{}
	if r.Method == "POST" {
		ButtonValidationOui := r.FormValue("ButtonOui")

		if ButtonValidationOui != "" {
			intid, err := strconv.Atoi(ButtonValidationOui)
			if err != nil {
				fmt.Println("error : ", err)
			}
			Delete(intid)
		}
	}
	if err == nil {
		Poste := PostByUser(cookie.Value)
		data.Poste = Poste
	} else {
		log.Println("vous n'etes pas connecter")
	}
	user := searchUserByUUID(cookie.Value)
	data.User = user
	tmpl, err := template.ParseFiles("../template/profile.html")
	if err != nil {
	}
	tmpl.ExecuteTemplate(w, "profile", data)
}
