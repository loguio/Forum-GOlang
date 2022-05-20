package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	_ "github.com/satori/go.uuid"
	uuid "github.com/satori/go.uuid"
)

var sessionStore = map[string]string{} // preferably uuid as key but username would work here

func main() {
	var err = error(nil)
	u1 := uuid.Must(uuid.NewV1(), err)
	fmt.Printf("UUIDv1: %s\n", u1)

	fileServer := http.FileServer(http.Dir("static/")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", handler)              // lance l'erreur 404 quand on est sur une URL pas utilisée
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
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not implemented")
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./sign in.html") // utilisation du fichier navPage.gohtml pour le template
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

func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./Log.html")
	if err != nil {
	}
	if r.Method == "POST" {
		cookie, err := r.Cookie("session-id")
		if err == nil {
			_, ok := sessionStore[cookie.Value]
			if ok {
				fmt.Fprintf(w, "You've already logged in.")
				return
			}
		} else if err != nil {
			cookie = &http.Cookie{
				Name: "session-id",
			}
		}

		UserName := r.FormValue("Username")
		Password := r.FormValue("password")
		if UserName != "" && Password != "" {
			user = User{Username: UserName, Password: Password, Email: "test"}
			loginSQL(user)

			cookie.Value = UserName
			code := getCode(cookie.Value)
			sessionStore[code] = UserName

			http.SetCookie(w, cookie)

			w.WriteHeader(http.StatusOK)

		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Credential not found")
			return
		}
	}
	tmpl.ExecuteTemplate(w, "Login", user)
}

func home(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./home.html")
	if err != nil {
	}
	tmpl.ExecuteTemplate(w, "home", user)
}
