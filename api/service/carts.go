package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"errors"
	"log"
)

type CartsService interface {
	ListCarts(ctx context.Context, userId string) (*models.Cart, error)
	AddCarts(ctx context.Context, userID, productName string, quantity int) (string, error)
	UpdateCarts(ctx context.Context, productName string, quantity int, args []string) (bool, error)
}

type cartsService struct {
	repository repository.Repository
	logger     log.Logger
}

func NewServiceCarts(repo repository.Repository, logger log.Logger) CartsService {
	return &cartsService{
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

// list cart from username
func (s cartsService) ListCarts(ctx context.Context, userId string) (*models.Cart, error) {
	fields := make(map[string]interface{})
	fields["username"] = userId
	cartData, err := s.repository.Get(ctx, &models.Cart{}, fields)
	if err != nil {
		return nil, err
	}
	cart := cartData.(*models.Cart)

	return cart, nil
}

// Create a cart with an product item into it
func (s cartsService) AddCarts(
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
		ProductID:  product.ID,
		Name:       product.Name,
		Quantity:   quantity,
		TotalPrice: totalPrice,
	}
	cart.CartItems = append(cart.CartItems, cartItem)
	cart.Quantity = quantity
	cart.Username = user.Username

	log.Println("CART IN ADD: ", cart)
	ok, err := s.repository.Create(ctx, &cart)
	if err != nil {
		return "", err
	}

	return ok, nil
}

// Update cart with new product item
func (s cartsService) UpdateCarts(ctx context.Context, productName string,
	quantity int, args []string) (bool, error) {

	var cartFields = make(map[string]interface{})
	cartFields["username"] = args[0]
	data, err := s.repository.Get(ctx, &models.Cart{}, cartFields)
	if err != nil {
		return false, err
	}
	cart := data.(*models.Cart)

	var errNoUserFound = errors.New("no cart found for user:" + cart.Username)
	if len(cart.Username) == 0 {
		return false, errNoUserFound
	}

	var productField = make(map[string]interface{})
	productField["name"] = productName
	dataProduct, err := s.repository.Get(ctx, &models.Product{}, productField)
	if err != nil {
		return false, err
	}

	product := dataProduct.(*models.Product)

	cartItemUpdated := models.CartItem{
		CartID:     cart.ID,
		ProductID:  product.ID,
		Name:       productName,
		Quantity:   quantity,
		TotalPrice: float64(quantity) * productPrice[productName],
	}

	var updateFields = make(map[string]interface{})
	for index, cartItem := range cart.CartItems {
		if cartItem.Name != cartItemUpdated.Name {
			_, err := s.repository.Create(ctx, &cartItemUpdated)
			if err != nil {
				return false, err
			}
			return true, nil
		} else {
			updateFields["quantity"] = 666
			updateFields["total_price"] = 666.666
			ok, err := s.repository.Update(ctx, &cart.CartItems[index], args[1], updateFields)
			if err != nil {
				return false, err
			}
			return ok, nil
		}
	}

	return true, nil
}
