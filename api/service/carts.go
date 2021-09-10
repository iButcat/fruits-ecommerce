package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"log"
)

type ServiceCarts interface {
	ListCarts(ctx context.Context) (models.Cart, error)
	AddCarts(ctx context.Context, cart models.Cart) (bool, error)
	UpdateCarts(ctx context.Context, id string) (bool, error)
	DeleteCarts(ctx context.Context, id string) (bool, error)
}

type serviceCarts struct {
	repository repository.Repository
	logger     log.Logger
}

func NewServiceCarts(repo repository.Repository, logger log.Logger) ServiceCarts {
	return &serviceCarts{
		repository: repo,
		logger:     logger,
	}
}

func (s serviceCarts) ListCarts(ctx context.Context) (models.Cart, error) {
	return models.Cart{}, nil
}

func (s serviceCarts) AddCarts(ctx context.Context, cart models.Cart) (bool, error) {
	return true, nil
}

func (s serviceCarts) UpdateCarts(ctx context.Context, id string) (bool, error) {
	return true, nil
}

func (s serviceCarts) DeleteCarts(ctx context.Context, id string) (bool, error) {
	return true, nil
}
