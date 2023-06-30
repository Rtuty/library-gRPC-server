package models

type Book struct {
	ID          string `json:"book_id"`
	Title       string `json:"book_name"`
	AuthorId    string `json:"book_author_id"`
	Description string `json:"book_description"`
}
