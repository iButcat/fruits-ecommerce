package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"ecommerce/utils"
	"errors"
	"log"

	"gorm.io/gorm"
)

type ServiceCarts interface {
	ListCarts(ctx context.Context, userId string) (*models.Cart, error)
	AddCarts(ctx context.Context, userID, productName string, quantity int) (string, error)
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

var productPrice = map[string]float64{
	"apple":   0.70,
	"bananas": 0.85,
	"oranges": 0.67,
	"pears":   0.85,
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
	ctx context.Context, userID, productName string, quantity int) (string, error) {
	var userParams = make(map[string]interface{})
	userParams["username"] = userID
	dataUser, err := s.repository.Get(ctx, &models.User{}, userParams)
	if err != nil {
		return "", err
	}
	user := dataUser.(*models.User)

	var cartParams = make(map[string]interface{})
	cartParams["name"] = productName

	price := productPrice[productName] * float64(quantity)

	cart := models.Cart{}
	product := &models.Product{
		Name:     productName,
		Quantity: quantity,
		Price:    price,
	}
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

	var errNoUserFound = errors.New("no cart found for user:" + cart.Username)
	if len(cart.Username) == 0 {
		return false, errNoUserFound
	}
	price := productPrice[name] * float64(quantity)
	cartUpdated := &models.Product{
		Model:    gorm.Model{ID: 5},
		Name:     name,
		Quantity: 18,
		Price:    price,
	}
	price = utils.CalculateDiscountBanana(quantity, cartUpdated.Price)

	var field = make(map[string]interface{})
	for _, value := range cart.Product {
		if value.ID == cartUpdated.ID {
			field["quantity"] = cartUpdated.Quantity
			field["price"] = price
			s.repository.Update(ctx, value, "5", field)
			break
		} else {
			cart.Product = append(cart.Product, cartUpdated)
			ok, err := s.repository.UpdateNested(ctx, &cart)
			if err != nil {
				return false, err
			}
			log.Println("HERE2")
			log.Println(ok)
		}
	}
	return true, nil
}
