package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"log"
)

type ServiceCarts interface {
	ListCarts(ctx context.Context, userId string) (*models.Cart, error)
	AddCarts(ctx context.Context, userID, productName string, quantity int, price float64) (string, error)
	UpdateCarts(ctx context.Context, id string, name string, quantity int) (bool, error)
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

func (s serviceCarts) AddCarts(
	ctx context.Context, userID, productName string, quantity int, price float64) (string, error) {
	var userParams = make(map[string]interface{})
	userParams["username"] = userID
	dataUser, err := s.repository.Get(ctx, &models.User{}, userParams)
	if err != nil {
		return "", err
	}
	user := dataUser.(*models.User)

	var cartParams = make(map[string]interface{})
	cartParams["name"] = productName

	cart := models.Cart{}
	product := &models.Product{
		Name:     productName,
		Quantity: quantity,
		Price:    price}
	cart.Product = append(cart.Product, product)
	cart.Quantity = quantity
	cart.Username = user.Username

	ok, err := s.repository.Create(ctx, &cart)
	if err != nil {
		return "", err
	}

	return ok, nil
}

func (s serviceCarts) UpdateCarts(ctx context.Context, id string, name string, quantity int) (bool, error) {
	var fields = make(map[string]interface{})
	fields["username"] = id
	data, err := s.repository.Get(ctx, &models.Cart{}, fields)
	if err != nil {
		return false, err
	}
	cart := data.(*models.Cart)
	cart.Product = append(cart.Product, &models.Product{Name: name, Quantity: quantity})
	ok, err := s.repository.UpdateNested(ctx, &cart, &models.Cart{})
	if err != nil {
		return false, err
	}
	return ok, nil
}
