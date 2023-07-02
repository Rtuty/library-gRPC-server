package transport

import (
	"context"
	"log"
	"modules/internal/db"
	"modules/internal/models"
	"modules/internal/transport/library"
)

type LibraryServer struct {
	Db db.Repository
	library.UnimplementedLibraryServer
}

func (s *LibraryServer) HandleAuthorCUD(ctx context.Context, ls *library.CUDAuthorRequest) (*library.DefaultResponse, error) {
	var a models.Author

	a.ID = ls.Author.ID
	a.Name = ls.Author.Name
	a.Country = ls.Author.Country

	if err := s.Db.AuthorMethodsHandler(ctx, ls.Operation, a); err != nil {
		log.Fatal("AuthorMethodsHandler error")
	}
	return &library.DefaultResponse{Result: "Done"}, nil
}

func (s *LibraryServer) HandleBookCUD(ctx context.Context, ls *library.CUDBookRequest) (*library.DefaultResponse, error) {
	var b models.Book

	b.ID = ls.Book.ID
	b.Title = ls.Book.Title
	b.Title = ls.Book.AuthorId
	b.Description = ls.Book.Description

	if err := s.Db.BookMethodsHandler(ctx, ls.Operation, b); err != nil {
		log.Fatal("BookMethodsHandler error")
	}
	return &library.DefaultResponse{Result: "Done"}, nil
}
