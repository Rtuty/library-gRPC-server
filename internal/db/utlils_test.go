package db

import (
	"modules/internal/models"
	"testing"
)

func TestConvertEntity(t *testing.T) {
	authors := []interface{}{
		models.Author{
			ID:      1,
			Name:    "Author 1",
			Country: "Country 1",
		},
		models.Author{
			ID:      2,
			Name:    "Author 2",
			Country: "Country 2",
		},
	}

	convertedAuthors, err := convertEntity[models.Author](authors)
	if err != nil {
		t.Errorf("convertEntity returned an error: %v", err)
	}

	if len(convertedAuthors) != 2 {
		t.Errorf("Expected 2 converted authors, got %d", len(convertedAuthors))
	}

	books := []interface{}{
		models.Book{
			ID:          1,
			Title:       "Title 1",
			AuthorId:    "AuthorId 1",
			Description: "Description 1",
		},
		models.Book{
			ID:          2,
			Title:       "Title 2",
			AuthorId:    "AuthorId 2",
			Description: "Description 2",
		},
	}

	convertedBooks, err := convertEntity[models.Book](books)
	if err != nil {
		t.Errorf("convertEntity returned an error: %v", err)
	}

	if len(convertedBooks) != 2 {
		t.Errorf("Expected 2 converted books, got %d", len(convertedBooks))
	}
}
