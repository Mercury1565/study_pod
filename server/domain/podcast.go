package domain

import (
	"context"
)

type Podcast struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"Description"`
	Url         string `json:"url"`
}

type PodcastRepository interface {
	Create(c context.Context, Podcast *Podcast) error
	FetchByID(c context.Context, id string) (*Podcast, error)
}

type PodcastUseCase interface {
	Create(c context.Context, Podcast *Podcast) error
	FetchByID(c context.Context, id string) (*Podcast, error)
}
