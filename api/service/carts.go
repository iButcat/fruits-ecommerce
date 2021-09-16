package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"log"
	"strconv"
)

type ServiceCarts interface {
	ListCarts(ctx context.Context, userId string) (*models.Cart, error)
	AddCarts(ctx context.Context, cart models.Cart, params ...string) (string, error)
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

func (s serviceCarts) AddCarts(ctx context.Context, cart models.Cart, params ...string) (string, error) {
	var userParams = make(map[string]interface{})
	userParams["username"] = params[1]
	dataUser, err := s.repository.Get(ctx, &models.User{}, userParams)
	if err != nil {
		return "", err
	}
	user := dataUser.(*models.User)

	var cartParams = make(map[string]interface{})
	cartParams["name"] = params[0]
	dataProduct, err := s.repository.Get(ctx, &models.Product{}, cartParams)
	if err != nil {
		return "", err
	}

	// convert interface to model
	product := dataProduct.(*models.Product)
	cart.Username = user.Username
	cart.Product = append(cart.Product, product)
	cart.Quantity, _ = strconv.Atoi(params[2])

	ok, err := s.repository.Create(ctx, &cart)
	if err != nil {
		return "", err
	}

	return ok, nil
}

func (s serviceCarts) UpdateCarts(ctx context.Context, id string) (bool, error) {
	var fields = make(map[string]interface{})
	fields["id"] = 1
	data, err := s.repository.Get(ctx, &models.Cart{}, fields)
	if err != nil {
		return false, err
	}
	cart := data.(*models.Cart)
	cart.Product = append(cart.Product, &models.Product{Name: "apple"})
	ok, err := s.repository.UpdateNested(ctx, &cart, &models.Cart{})
	if err != nil {
		return false, err
	}
	return ok, nil
}
