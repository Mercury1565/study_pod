package domain

import (
	"context"
)

type Instance struct {
	ID          string `json:"id"`
	UserID    string `json:"user_id"`
	BookID          string `json:"book_id"`
	ChapterID *string `json:"chapter_id"`
	PodcastID *string `json:"podcast_id"`
}

type InstanceRepository interface {
	Create(c context.Context, Instance *Instance) error
	FetchByID(c context.Context, id string) (*Instance, error)
	FetchByUserID(c context.Context, userId string) ([]Instance, error)
}

type InstanceUseCase interface {
	Create(c context.Context, Instance *Instance) error
	FetchByID(c context.Context, id string) (*Instance, error)
	FetchByUserID(c context.Context, userId string) ([]Instance, error)
}
