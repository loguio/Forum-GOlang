package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	_ "github.com/satori/go.uuid"
)

var sessionStore = map[string]string{} // preferably uuid as key but username would work here
type data struct {
	Posts     []Post
	Connected bool
}

func main() {
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
	case "/logout":
		Logout(w, r)
	case "/profile":
		profile(w,r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not implemented")
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./template/signIn.html") // utilisation du fichier navPage.gohtml pour le template
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

var secret = "secret"

func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./template/log.html")
	if err != nil {
	}
	if r.Method == "POST" {
		cookie, err := r.Cookie("session-id")
		if err == nil {
			sessionStore["session"] = cookie.Value
			fmt.Println("previous cookie value : " + cookie.Value)
			fmt.Println("You've already logged in.")
		} else if err != nil {
			fmt.Println("cookie not found")
			cookie = &http.Cookie{
				Name: "session-id",
			}

			UserName := r.FormValue("Username")
			Password := r.FormValue("password")
			if UserName != "" && Password != "" {
				user = User{Username: UserName, Password: Password}

				connected := loginSQL(user)
				if connected == true {
					sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
					defer sqliteDatabase.Close()
					user = searchUser(sqliteDatabase, UserName)

					cookie.Value = user.UUID
					cookie.MaxAge = 604800
					sessionStore["session"] = user.UUID

					http.SetCookie(w, cookie)

					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Println("entrée vide")
			}
		}
	}
	fmt.Println("---------------------------------")
	tmpl.ExecuteTemplate(w, "Login", user)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./Log.html")
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

func home(w http.ResponseWriter, r *http.Request) {
	post := Post{}
	var data data
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()
	tri := []Post{}
	tmpl, err := template.ParseFiles("./template/home.html")
	if err != nil {
	}
	cookie, err := r.Cookie("session-id")
	if err == nil {
		data.Connected = true
		sessionStore["session"] = cookie.Value
		if r.Method == "POST" {
			Name := r.FormValue("Name")
			Contentpost := r.FormValue("Contentpost")
			Categorie := r.FormValue("Categorie")
			buttonSelect := r.FormValue("buttonCategorie")
			if buttonSelect != "" {
				fmt.Println("buttonSelect : " + buttonSelect)
				tri = triPost(sqliteDatabase, buttonSelect)
				
			} else {
				tri = postDB()
				post = Post{Name: Name, Contentpost: Contentpost, Categorie: Categorie}
				addpost(sqliteDatabase, post.Name, post.Contentpost, post.Categorie)
			}

		}
	} else {
		data.Connected = false
		buttonSelect := r.FormValue("buttonCategorie")
		if buttonSelect != "" {
			fmt.Println("buttonSelect : " + buttonSelect)
			tri = triPost(sqliteDatabase, buttonSelect)
			
		} else {
			tri = postDB()
		}
	}
	if tri != nil {
		data.Posts = tri
	} else {
		data.Posts = postDB()
	}
	

	tmpl.ExecuteTemplate(w, "home", data)
}

func profile(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("./template/profile.html")
	if err != nil {
	}
	tmpl.ExecuteTemplate(w, "profile", user)
}
