package models

type Author struct {
	ID      int64  `json:"author_id"`
	Name    string `json:"author_name"`
	Country string `json:"author_country"`
}
