package controller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	// internal pkg

	"ecommerce/service"
	"encoding/json"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type CartController interface {
	Add(ctx *gin.Context)
	List(ctx *gin.Context)
}

type cartController struct {
	service service.CartsService
	logger  log.Logger
}

func NewCartController(service service.CartsService, logger log.Logger) CartController {
	return &cartController{
		service: service,
		logger:  logger,
	}
}

var (
	errNoProductName      = errors.New("err no product name provided")
	errNoQuantityProvided = errors.New("err no quantity provided")
)

func (c cartController) Add(ctx *gin.Context) {
	var addRequestBody struct {
		ProductName string `json:"product_name"`
		Quantity    int    `json:"quantity"`
	}

	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	json.Unmarshal(data, &addRequestBody)

	quantity := addRequestBody.Quantity
	productName := addRequestBody.ProductName
	if len(productName) == 0 {
		ctx.JSON(400, gin.H{"error": errNoProductName.Error()})
		return
	} else if quantity == 0 {
		ctx.JSON(400, gin.H{"error": errNoQuantityProvided.Error()})
		return
	}

	claims := jwt.ExtractClaims(ctx)

	userID := claims["id"].(string)

	cart, err := c.service.ListCarts(ctx, userID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	var args []string
	args = append(args,
		claims["id"].(string),
		fmt.Sprint(cart.ID),
		addRequestBody.ProductName)
	for _, cartItem := range cart.CartItems {
		if productName == cartItem.Name {
			success, err := c.service.UpdateCarts(ctx,
				addRequestBody.ProductName, addRequestBody.Quantity, args)
			if err != nil {
				log.Println(err)
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, gin.H{"sucess": success})
		} else if len(cartItem.Name) != 0 {
			_, err := c.service.UpdateCarts(ctx, productName, addRequestBody.Quantity, args)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
		}
	}

	success, err := c.service.AddCarts(ctx, userID, productName, quantity)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": success})
}

func (c cartController) List(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	userID := claims["id"].(string)
	cart, err := c.service.ListCarts(ctx, userID)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"cart": cart})
}
