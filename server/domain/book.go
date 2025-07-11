package domain

import (
	"context"
)

type Book struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
}

type BookRepository interface {
	Create(c context.Context, Book *Book) error
	FetchByUserID(c context.Context, userID string) ([]Book, error)
}

type BookUseCase interface {
	Create(c context.Context, Book *Book) error
	FetchByUserID(c context.Context, userID string) ([]Book, error)
}
