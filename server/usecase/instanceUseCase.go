package usecase

import (
	"Clean_Architecture/domain"
	"context"
	"time"
)

type InstanceUsecase struct {
	InstanceRepository domain.InstanceRepository
	contextTimeout time.Duration
}

func NewInstanceUsecase(InstanceRepository domain.InstanceRepository, timeout time.Duration) domain.InstanceUseCase {
	return &InstanceUsecase{
		InstanceRepository: InstanceRepository,
		contextTimeout: timeout,
	}
}

func (instanceUC *InstanceUsecase) Create(c context.Context, instance *domain.Instance) error {
	ctx, cancel := context.WithTimeout(c, instanceUC.contextTimeout)
	defer cancel()

	return instanceUC.InstanceRepository.Create(ctx, instance)
}

func (instanceUC *InstanceUsecase) FetchByID(c context.Context, instanceId string) (*domain.Instance, error) {
	ctx, cancel := context.WithTimeout(c, instanceUC.contextTimeout)
	defer cancel()

	return instanceUC.InstanceRepository.FetchByID(ctx, instanceId)
}

func (instanceUC *InstanceUsecase) FetchByUserID(c context.Context, instanceId string) ([]domain.Instance, error) {
	ctx, cancel := context.WithTimeout(c, instanceUC.contextTimeout)
	defer cancel()

	return instanceUC.InstanceRepository.FetchByUserID(ctx, instanceId)
}
