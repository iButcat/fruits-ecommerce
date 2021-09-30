package controller

import (
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
	Update(ctx *gin.Context)
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

func (c cartController) Add(ctx *gin.Context) {
	var addRequestBody struct {
		ProductID   string `json:"product_id"`
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

	claims := jwt.ExtractClaims(ctx)

	userID := claims["id"].(string)

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

func (c cartController) Update(ctx *gin.Context) {
	var updateBodyRequest struct {
		CartID      string `json:"cart_id"`
		ProductName string `json:"product_name"`
		Quantity    int    `json:"quantity"`
	}

	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	json.Unmarshal(data, &updateBodyRequest)

	claims := jwt.ExtractClaims(ctx)

	var args []string
	args = append(args,
		claims["id"].(string),
		updateBodyRequest.CartID,
		updateBodyRequest.ProductName)

	success, err := c.service.UpdateCarts(ctx, updateBodyRequest.Quantity, args)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": success})
}
