package transport

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"modules/internal/db"
	"modules/internal/models"
	"modules/internal/transport/library"
)

type LibraryServer struct {
	Db db.Repository
	library.UnimplementedLibraryServer
}

/*
CUD:
	C-create,
	U-update,
	D-delete
DataBase sql functionality with local entities
*/

// HandleAuthorCUD обрабатывает запрос от клиента на создание, удаление и обновление сущности автор
func (s *LibraryServer) HandleAuthorCUD(ctx context.Context, ls *library.CUDAuthorRequest) (*library.DefaultResponse, error) {
	var a models.Author

	a.ID = ls.Author.ID
	a.Name = ls.Author.Name
	a.Country = ls.Author.Country

	if err := s.Db.AuthorMethodsHandler(ctx, ls.Operation, a); err != nil {
		return nil, status.Errorf(codes.Aborted, "HandleAuthorCUD error: %s", err)
	}
	return &library.DefaultResponse{Result: "Done"}, nil
}

// HandleBookCUD обрабатывает запрос от клиента на создание, удаление и обновление сущности книга
func (s *LibraryServer) HandleBookCUD(ctx context.Context, ls *library.CUDBookRequest) (*library.DefaultResponse, error) {
	var b models.Book

	b.ID = ls.Book.ID
	b.Title = ls.Book.Title
	b.Title = ls.Book.AuthorId
	b.Description = ls.Book.Description

	if err := s.Db.BookMethodsHandler(ctx, ls.Operation, b); err != nil {
		return nil, status.Errorf(codes.Aborted, "BookMethodsHandler error: %s", err)
	}
	return &library.DefaultResponse{Result: "Done"}, nil
}

// GetAllAuthors возращает список авторов из базы данных
func (s *LibraryServer) GetAllAuthors(ctx context.Context, ls *library.DefaultRequest) (*library.Authors, error) {
	a, err := s.Db.GetAllAuthors(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "GetAllAuthors error: %s", err)
	}

	res := &library.Authors{
		Authors: []*library.Author{},
	}

	for _, v := range a {
		res.Authors = append(res.Authors,
			&library.Author{
				ID:      v.ID,
				Name:    v.Name,
				Country: v.Country,
			})
	}

	return res, nil
}

// GetAllBooks возращает список книг из базы данных
func (s *LibraryServer) GetAllBooks(ctx context.Context, ls *library.DefaultRequest) (*library.Books, error) {
	b, err := s.Db.GetAllBooks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "GetAllBooks error: %s", err)
	}

	res := &library.Books{
		Books: []*library.Book{},
	}

	for _, v := range b {
		res.Books = append(res.Books,
			&library.Book{
				ID:          v.ID,
				Title:       v.Title,
				AuthorId:    v.AuthorId,
				Description: v.Description,
			})
	}

	return res, nil
}

func (s *LibraryServer) GetAuthorById(ctx context.Context, ls *library.IdRequest) (*library.Author, error) {
	a, err := s.Db.GetAuthorById(ctx, ls.Id)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "GetAuthorById error: %s", err)
	}

	res := &library.Author{
		ID:      a.ID,
		Name:    a.Name,
		Country: a.Country,
	}

	return res, nil
}

func (s *LibraryServer) GetBookById(ctx context.Context, ls *library.IdRequest) (*library.Book, error) {
	b, err := s.Db.GetBookById(ctx, ls.Id)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "GetBookById error: %s", err)
	}

	res := &library.Book{
		ID:          b.ID,
		Title:       b.Title,
		AuthorId:    b.AuthorId,
		Description: b.Description,
	}

	return res, nil
}

func (s *LibraryServer) GetAuthorsByBookName(ctx context.Context, ls *library.Book) (*library.Authors, error) {
	a, err := s.Db.GetAuthorsByBookName(ctx, ls.Title)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "GetAuthorsByBookName error: %s", err)
	}

	res := &library.Authors{
		Authors: []*library.Author{},
	}

	for _, v := range a {
		res.Authors = append(res.Authors,
			&library.Author{
				ID:      v.ID,
				Name:    v.Name,
				Country: v.Country,
			})
	}

	return res, nil
}

func (s *LibraryServer) GetBooksByAuthorId(ctx context.Context, ls *library.IdRequest) (*library.Books, error) {
	b, err := s.Db.GetBooksByAuthorId(ctx, ls.Id)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "GetBooksByAuthorId error: %s", err)
	}

	res := &library.Books{
		Books: []*library.Book{},
	}

	for _, v := range b {
		res.Books = append(res.Books,
			&library.Book{
				ID:          v.ID,
				Title:       v.Title,
				AuthorId:    v.AuthorId,
				Description: v.Description,
			})
	}

	return res, nil
}
