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

func (s *LibraryServer) HandlerAuthorCUD(ctx context.Context, ls *library.CUDAuthorRequest) (*library.CUDResponse, error) {
	var a models.Author

	a.ID = ls.Author.ID
	a.Name = ls.Author.Name
	a.Country = ls.Author.Country

	if err := s.Db.AuthorMethodsHandler(ctx, ls.Operation, a); err != nil {
		log.Fatal("AuthorMethodsHandler error")
	}
	return &library.CUDResponse{Result: "Done"}, nil
}
