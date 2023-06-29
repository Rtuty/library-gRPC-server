package models

type Book struct {
	ID     string `json:"book_id"`
	Name   string `json:"book_name"`
	Author Author
}
