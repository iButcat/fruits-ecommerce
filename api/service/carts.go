package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"ecommerce/utils"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

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
	cartData, err := s.repository.Get(ctx, &models.Cart{}, fields)
	if err != nil {
		return nil, err
	}
	cart := cartData.(*models.Cart)

	cartItemData, err := s.repository.FindAll(ctx, &[]models.CartItem{}, fmt.Sprint("cart_id = ", cart.ID))
	if err != nil {
		return nil, err
	}
	cartItem := cartItemData.(*[]models.CartItem)
	log.Println("CART ITEM: ", cartItem)

	cart.CartItems = append(cart.CartItems, *cartItem...)

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

	totalPrice := productPrice[productName] * float64(quantity)

	cart := models.Cart{}

	var fields = make(map[string]interface{})
	fields["name"] = productName
	dataProduct, err := s.repository.Get(ctx, &models.Product{}, fields)
	if err != nil {
		return "err while querying product: ", err
	}
	product := dataProduct.(*models.Product)

	cartItem := models.CartItem{
		CartID:     cart.ID,
		ProductID:  product.ID,
		Name:       product.Name,
		Quantity:   quantity,
		TotalPrice: totalPrice,
	}
	cart.CartItems = append(cart.CartItems, cartItem)
	cart.Quantity = quantity
	cart.Username = user.Username

	ok, err := s.repository.Create(ctx, &cart)
	if err != nil {
		return "", err
	}

	return ok, nil
}

func (s serviceCarts) UpdateCarts(ctx context.Context, id string, cartID string, quantity int) (bool, error) {
	var fields = make(map[string]interface{})
	fields["username"] = id
	data, err := s.repository.Get(ctx, &models.Cart{}, fields)
	if err != nil {
		return false, err
	}
	cart := data.(*models.Cart)
	log.Println("CART GET IN UPDATE", cart)

	var errNoUserFound = errors.New("no cart found for user:" + cart.Username)
	if len(cart.Username) == 0 {
		return false, errNoUserFound
	}
	price := productPrice["apple"] * float64(quantity)
	u64, err := strconv.ParseUint(cartID, 10, 32)
	if err != nil {
		log.Println(err)
	}
	productUpdated := models.Product{
		Model: gorm.Model{ID: uint(u64)},
		Name:  "bananas",
		Price: price,
	}
	cartUpdated := models.Cart{
		Model:      gorm.Model{ID: cart.ID},
		Quantity:   3,
		TotalPrice: 1000,
		Username:   cart.Username,
	}
	price = utils.CalculateDiscountBanana(quantity, productUpdated.Price)

	var field = make(map[string]interface{})
	for _, value := range cart.CartItems {
		if value.Name == productUpdated.Name {
			log.Println("BUG HERE ONE")
			field["price"] = price
			s.repository.Update(ctx, &value, fmt.Sprint(value.ID), field)
			break
		} else {
			log.Println("BUG HERE TWO")
			cartUpdated.CartItems = append(cartUpdated.CartItems, cart.CartItems...)
			ok, err := s.repository.AppendNested(ctx, &cart, []models.CartItem{
				{
					Model: gorm.Model{
						ID:        6,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
					CartID:     1,
					ProductID:  4,
					Name:       "oranges",
					Quantity:   12121212,
					TotalPrice: 1111212.34444,
				},
			})
			if err != nil {
				return false, err
			}
			log.Println(ok)
		}
	}
	return true, nil
}
