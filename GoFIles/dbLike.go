package main

import (
	"database/sql"
	"log"
	"strings"
)

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func dbLike(id int, UUID string) {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()
	// fmt.Println(searchLikePost(id))
	like := searchLikePost(id)
	var there bool
	IDLike := strings.Split(like, " ")
	var i int
	for i = 0; i < len(IDLike); i++ {
		if IDLike[i] == UUID {
			there = true
			break
		}
	}
	if there {
		IDLike = remove(IDLike, i)
		like = strings.Join(IDLike, " ")
		addLike := `UPDATE TablePost SET Like = ? WHERE id = ?`
		statement, err := db.Prepare(addLike)
		if err != nil {
			log.Println(err.Error())
		}
		_, err = statement.Exec(like, id)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("like removed")
	} else {
		addLike := `UPDATE TablePost SET Like = ? WHERE id = ?`
		statement, err := db.Prepare(addLike)
		if err != nil {
			log.Println(err.Error())
		}
		_, err = statement.Exec(like+" "+UUID, id)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("like added")
	}
}
