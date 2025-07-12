package domain

import (
	"context"
)

type Chapter struct {
	ID          string `json:"id"`
	BookID          string `json:"book_id"`
	Chapter   string `json:"chapter"`
	Title       string `json:"title"`
	Description string `json:"Description"`
}

type ChapterRepository interface {
	Create(c context.Context, Chapter *Chapter) error
	FetchByID(c context.Context, id string) (*Chapter, error)
}

type ChapterUseCase interface {
	Create(c context.Context, Chapter *Chapter) error
	FetchByID(c context.Context, id string) (*Chapter, error)
}
