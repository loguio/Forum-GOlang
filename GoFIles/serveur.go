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

	http.HandleFunc("/", handler)              // Handler pour les requêtes HTTP
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path { //pour chaque URL possible sur le serveur web
	case "/Login":
		Login(w, r) //lance la fonction Login
	case "/Signin":
		Signin(w, r) //lance la fonction Signin
	case "/home":
		home(w, r)
	case "/profile":
		profile(w, r)
	case "/logout":
		Logout(w, r) //lance la fonction Logout
	default:
		w.WriteHeader(http.StatusNotFound) // lance l'erreur 404 quand on est sur une URL pas utilisée
		erreur404(w)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../template/log.html") //charge le template
	if err != nil {
		erreur500(w)
	}
	cookie, err := r.Cookie("session-id") //récupere le cookie session-id de la page
	if err == nil {                       // si il n'y en a pas
		fmt.Println("cookie value before logout : " + cookie.Value)
		cookie.MaxAge = -1        //supprime le cookie
		http.SetCookie(w, cookie) //envoie le cookie
		log.Println("You've successfully logged out.")

	} else {
		log.Println("You're not logged in.")
	}
	tmpl.ExecuteTemplate(w, "Login", nil) //execute le template

}

func erreur500(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("../template/500.html") //charge le template
	if err != nil {
		fmt.Fprintf(w, "Erreur 500")
	}
	tmpl.ExecuteTemplate(w, "500", nil) //execute le template
}

func erreur404(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("../template/404.html") //charge le template
	if err != nil {
		fmt.Fprintf(w, "Erreur 404")
	}
	tmpl.ExecuteTemplate(w, "404", nil) //execute le template
}

func profile(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("../template/profile.html") //charge le template
	if err != nil {
		erreur500(w)
	}
	cookie, err := r.Cookie("session-id")
	data := DataProfile{}
	if r.Method == "POST" {
		ButtonValidationOui := r.FormValue("ButtonOui")
		IDButtonLike := r.FormValue("buttonLike")
		if ButtonValidationOui != "" {
			intid, err := strconv.Atoi(ButtonValidationOui)
			if err != nil {
				fmt.Println("error : ", err)
			}
			Delete(intid)
		} else if IDButtonLike != "" {
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
		}
	}
	if err == nil {
		Poste := PostByUser(cookie.Value)
		data.Poste = Poste
		user = searchUserByUUID(cookie.Value)
		data.User = user
	} else {
		log.Println("vous n'etes pas connecter")
	}

	tmpl.ExecuteTemplate(w, "profile", data)
}
