package model

type Book struct {
	Id        int    `json: "id" form: "id"`
	Title     string `json: "title" form: "title"`
	Author    string `json: "author" form: "author"`
	Publisher string `json: "publisher" form: "publisher"`
}
