package main

type data struct {
	Posts     []Post
	Connected bool
}

type User struct {
	Username string
	Password string
	Email    string
	UUID     string
}

type Post struct {
	Name        string
	Contentpost string
	Categorie   string
	Like        int
	UUID        string
	ID          int
	Img			bool
}

type DataProfile struct {
	User		User
	Poste 		[]Post
}