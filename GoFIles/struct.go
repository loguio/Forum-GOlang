package main

type data struct { //création d'une structure a envoyer sur la page home
	Posts     []Post
	Connected bool
}

type User struct { //struct pour les utilisateurs
	Username string
	Password string
	Email    string
	UUID     string
}

type Post struct { //struct pour les posts
	Name        string
	Contentpost string
	Categorie   string
	Like        int
	UUID        string
	ID          int
	Comment     []Comment
}

type Comment struct { //struct pour les commentaires
	ID      int
	Comment string
	UUID    string
	Like    string
	IDPost  int
}

type DataProfile struct {
	User  User
	Poste []Post
}
