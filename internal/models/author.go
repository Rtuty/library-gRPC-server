package models

type Author struct {
	ID      string `json:"author_id"`
	Name    string `json:"author_name"`
	Country string `json:"author_country"`
}
