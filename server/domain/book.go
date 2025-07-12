package domain

import (
	"context"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type BookRepository interface {
	Create(c context.Context, Book *Book) error
	FetchByID(c context.Context, id string) (*Book, error)
}

type BookUseCase interface {
	Create(c context.Context, Book *Book) error
	FetchByID(c context.Context, id string) (*Book, error)
}
