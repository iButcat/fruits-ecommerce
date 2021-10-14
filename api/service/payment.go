package service

import (
	"context"
	"log"

	"ecommerce/models"
	"ecommerce/repository"
)

type ServicePayment interface {
	CreatePayment(ctx context.Context, userID, cartID string) (bool, error)
	UpdatePayment(ctx context.Context, amount float64) (bool, error)
}

type servicePayment struct {
	repository repository.Repository
	logger     log.Logger
}

func NewServicePayment(repo repository.Repository, logger log.Logger) ServicePayment {
	return &servicePayment{
		repository: repo,
		logger:     logger,
	}
}

func (s servicePayment) CreatePayment(ctx context.Context, userID, cartID string) (bool, error) {
	payment := new(models.Payment)
	payment.Username = userID
	payment.CartID = cartID

	var fields = make(map[string]interface{})
	fields["id"] = cartID
	dataCart, err := s.repository.Get(ctx, &models.Cart{}, fields)
	if err != nil {
		return false, err
	}
	cart := dataCart.(*models.Cart)

	payment.Amount = cart.TotalPrice

	_, err = s.repository.Create(ctx, &payment)
	if err != nil {
		return false, err
	}

	// delete cart after when is paid
	deleted, err := s.repository.Delete(ctx, &models.Cart{}, cartID)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

func (s servicePayment) UpdatePayment(ctx context.Context, amount float64) (bool, error) {
	return true, nil
}
