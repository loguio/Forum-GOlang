package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}
	tmpl, err := template.ParseFiles("../template/log.html")
	if err != nil {
	}
	if r.Method == "POST" {
		cookie, err := r.Cookie("session-id")
		if err == nil {
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
					user = searchUser(UserName)

					cookie.Value = user.UUID
					cookie.MaxAge = 604800

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
