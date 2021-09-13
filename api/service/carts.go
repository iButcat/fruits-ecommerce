package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"log"
)

type ServiceCarts interface {
	ListCarts(ctx context.Context, userId string) (*models.Cart, error)
	AddCarts(ctx context.Context, cart models.Cart, userId string) (string, error)
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

func (s serviceCarts) ListCarts(ctx context.Context, userId string) (*models.Cart, error) {
	fields := make(map[string]interface{})
	fields["user_id"] = userId
	data, err := s.repository.Get(ctx, &models.Cart{}, fields)
	if err != nil {
		return nil, err
	}
	cart := data.(*models.Cart)
	return cart, nil
}

func (s serviceCarts) AddCarts(ctx context.Context, cart models.Cart, userId string) (string, error) {
	var fields = make(map[string]interface{})
	fields["id"] = userId

	data, err := s.repository.Get(ctx, &models.User{}, fields)
	if err != nil {
		return "", err
	}

	user := data.(*models.User)
	cart.User = user

	created, err := s.repository.Create(ctx, &cart)
	if err != nil {
		return "", err
	}
	return created, nil
}

func (s serviceCarts) UpdateCarts(ctx context.Context, id string) (bool, error) {
	return true, nil
}

func (s serviceCarts) DeleteCarts(ctx context.Context, id string) (bool, error) {
	return true, nil
}
