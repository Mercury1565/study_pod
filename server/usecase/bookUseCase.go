package usecase

import (
	"Clean_Architecture/domain"
	"context"
	"time"
)

type BookUsecase struct {
	BookRepository domain.BookRepository
	contextTimeout time.Duration
}

func NewBookUsecase(BookRepository domain.BookRepository, timeout time.Duration) domain.BookUseCase {
	return &BookUsecase{
		BookRepository: BookRepository,
		contextTimeout: timeout,
	}
}

func (bookUC *BookUsecase) Create(c context.Context, book *domain.Book) error {
	ctx, cancel := context.WithTimeout(c, bookUC.contextTimeout)
	defer cancel()

	return bookUC.BookRepository.Create(ctx, book)
}

func (bookUC *BookUsecase) FetchByUserID(c context.Context, bookID string) ([]domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, bookUC.contextTimeout)
	defer cancel()

	return bookUC.BookRepository.FetchByUserID(ctx, bookID)
}
