package handlers

import (
	"context"
	"media-devoted/repositories"
	"media-devoted/types"

	"github.com/google/uuid"
)

type RocketHandler interface {
	GetRockets(ctx context.Context) (*[]types.Rocket, error)
	GetRocket(ctx context.Context, id *uuid.UUID) (*types.Rocket, error)
	AddRocket(ctx context.Context, rocket *types.Rocket) error
	UpdateRocket(ctx context.Context, rocket *types.Rocket) error
	DeleteRocket(ctx context.Context, id uuid.UUID) error
}

type RocketHandlerImpl struct {
	RocketRepository repositories.RocketRepository
}

func RocketHandlerInstance() RocketHandler {
	return &RocketHandlerImpl{
		RocketRepository: repositories.RocketRepositoryInstance(),
	}
}

func (h *RocketHandlerImpl) GetRockets(ctx context.Context) (*[]types.Rocket, error) {
	rockets, err := h.RocketRepository.GetRockets(ctx)
	if err != nil {
		return rockets, err
	}

	return rockets, nil
}

func (h *RocketHandlerImpl) GetRocket(ctx context.Context, id *uuid.UUID) (*types.Rocket, error) {
	rocket, err := h.RocketRepository.GetRocket(ctx, id)
	if err != nil {
		return rocket, err
	}
	return rocket, nil
}

func (h *RocketHandlerImpl) AddRocket(ctx context.Context, rocket *types.Rocket) error {
	id := uuid.New()

	rocket.Id = id
	err := h.RocketRepository.AddRocket(ctx, rocket)
	if err != nil {
		return err
	}

	return nil
}

func (h *RocketHandlerImpl) UpdateRocket(ctx context.Context, rocket *types.Rocket) error {
	updateErr := h.RocketRepository.UpdateRocket(ctx, rocket)
	if updateErr != nil {
		return updateErr
	}
	return updateErr
}

func (h *RocketHandlerImpl) DeleteRocket(ctx context.Context, id uuid.UUID) error {
	deleteErr := h.RocketRepository.DeleteRocket(ctx, id)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
