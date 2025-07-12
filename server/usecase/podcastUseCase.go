package usecase

import (
	"Clean_Architecture/domain"
	"context"
	"time"
)

type PodcastUsecase struct {
	PodcastRepository domain.PodcastRepository
	contextTimeout time.Duration
}

func NewPodcastUsecase(PodcastRepository domain.PodcastRepository, timeout time.Duration) domain.PodcastUseCase {
	return &PodcastUsecase{
		PodcastRepository: PodcastRepository,
		contextTimeout: timeout,
	}
}

func (podcastUC *PodcastUsecase) Create(c context.Context, podcast *domain.Podcast) error {
	ctx, cancel := context.WithTimeout(c, podcastUC.contextTimeout)
	defer cancel()

	return podcastUC.PodcastRepository.Create(ctx, podcast)
}

func (podcastUC *PodcastUsecase) FetchByID(c context.Context, podcastId string) (*domain.Podcast, error) {
	ctx, cancel := context.WithTimeout(c, podcastUC.contextTimeout)
	defer cancel()

	return podcastUC.PodcastRepository.FetchByID(ctx, podcastId)
}
