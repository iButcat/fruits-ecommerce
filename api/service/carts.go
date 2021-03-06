package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"ecommerce/utils"
	"errors"
	"fmt"
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
	"apples":  0.70,
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

	totalPriceDiscount := utils.CalculateDiscountApples(productName, quantity, totalPrice)
	if totalPriceDiscount != 0 {
		totalPrice = totalPriceDiscount
	}
	log.Println("total price after discount: ", totalPriceDiscount)

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
	cart.TotalPrice = cartItem.TotalPrice
	ok, err := s.repository.Create(ctx, &cart)
	if err != nil {
		return "", err
	}

	return ok, nil
}

// custom error for our update logic
var (
	errMissingArgs        = errors.New("err missing args for update carts")
	errNoProductNameFound = errors.New("no product name has been submitted")
	errNoUserFound        = errors.New("no cart found for user")
)

// Update cart with new product item
func (s cartsService) UpdateCarts(ctx context.Context, productName string,
	quantity int, args []string) (bool, error) {

	if len(args) != 3 {
		return false, errMissingArgs
	}

	var cartFields = make(map[string]interface{})
	cartFields["username"] = args[0]
	data, err := s.repository.Get(ctx, &models.Cart{}, cartFields)
	if err != nil {
		return false, err
	}
	cart := data.(*models.Cart)
	if len(cart.Username) == 0 {
		return false, errNoUserFound
	}

	if len(productName) == 0 {
		return false, errNoProductNameFound
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
		if cartItem.Name == cartItemUpdated.Name {
			updateFields["quantity"] = quantity
			updateFields["total_price"] = cartItemUpdated.TotalPrice
			_, err := s.repository.Update(ctx, &cart.CartItems[index],
				fmt.Sprint(cart.CartItems[index].ID), updateFields)
			if err != nil {
				return false, err
			}
			var updateFieldTotalPrice = make(map[string]interface{})
			updateFieldTotalPrice["total_price"] = utils.CalculateTotalPriceCart(*cart)
			success, err := s.repository.Update(ctx, &cart, fmt.Sprint(cart.ID), updateFieldTotalPrice)
			if err != nil {
				return false, err
			}
			log.Println(success)
			log.Println("Done")
			return true, nil
		}
	}
	success, err := s.repository.Create(ctx, &cartItemUpdated)
	if err != nil {
		return false, err
	}
	log.Println(success)

	var updateFieldTotalPrice = make(map[string]interface{})
	updateFieldTotalPrice["total_price"] = utils.CalculateTotalPriceCart(*cart)
	_, err = s.repository.Update(ctx, &cart, fmt.Sprint(cart.ID), updateFieldTotalPrice)
	if err != nil {
		return false, err
	}

	return true, nil
}
