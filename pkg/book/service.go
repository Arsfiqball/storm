package book

import "context"

type BookService struct {
	repo *BookRepository
}

func NewBookService(repo *BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (svc *BookService) CreateOne(ctx context.Context, payload PayloadBook) (EntityBook, error) {
	return svc.repo.CreateOne(ctx, payload)
}
