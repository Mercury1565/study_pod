package usecase

import (
	"Clean_Architecture/domain"
	"context"
	"time"
)

type ChapterUsecase struct {
	ChapterRepository domain.ChapterRepository
	contextTimeout time.Duration
}

func NewChapterUsecase(ChapterRepository domain.ChapterRepository, timeout time.Duration) domain.ChapterUseCase {
	return &ChapterUsecase{
		ChapterRepository: ChapterRepository,
		contextTimeout: timeout,
	}
}

func (chapterUC *ChapterUsecase) Create(c context.Context, chapter *domain.Chapter) error {
	ctx, cancel := context.WithTimeout(c, chapterUC.contextTimeout)
	defer cancel()

	return chapterUC.ChapterRepository.Create(ctx, chapter)
}

func (chapterUC *ChapterUsecase) FetchByID(c context.Context, chapterId string) (*domain.Chapter, error) {
	ctx, cancel := context.WithTimeout(c, chapterUC.contextTimeout)
	defer cancel()

	return chapterUC.ChapterRepository.FetchByID(ctx, chapterId)
}
