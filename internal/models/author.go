package models

type Author struct {
	ID    string `json:"author_id"`
	Name  string `json:"author_name"`
	Books []Book `json:"author_books"`
}
