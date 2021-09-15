package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"log"
)

type ServiceCarts interface {
	ListCarts(ctx context.Context, userId string) (*models.Cart, error)
	AddCarts(ctx context.Context, cart models.Cart, productName, username string) (string, error)
	UpdateCarts(ctx context.Context, id string) (bool, error)
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
	fields["username"] = userId
	data, err := s.repository.Get(ctx, &models.Cart{}, fields)
	if err != nil {
		return nil, err
	}
	cart := data.(*models.Cart)
	return cart, nil
}

func (s serviceCarts) AddCarts(ctx context.Context, cart models.Cart, productName, username string) (string, error) {
	dataUser, err := s.repository.Get(ctx, &models.User{}, map[string]interface{}{"username": username})
	if err != nil {
		return "", err
	}
	user := dataUser.(*models.User)

	var fields = map[string]interface{}{"name": productName}
	dataProduct, err := s.repository.Get(ctx, &models.Product{},
		fields)
	if err != nil {
		return "", err
	}
	product := dataProduct.(*models.Product)
	log.Print("PRODUCT: ", product)

	cart.Username = user.Username
	cart.Product = append(cart.Product, product)

	ok, err := s.repository.Create(ctx, &cart)
	if err != nil {
		return "", err
	}

	return ok, nil
}

func (s serviceCarts) UpdateCarts(ctx context.Context, id string) (bool, error) {
	return true, nil
}
