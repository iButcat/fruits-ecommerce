package controller

import (
	"ecommerce/service"
	"encoding/json"
	"io/ioutil"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	CreatePayment(ctx *gin.Context)
	UpdatePayment(ctx *gin.Context)
}

type paymentController struct {
	service service.ServicePayment
	logger  log.Logger
}

func NewPaymentController(service service.ServicePayment, logger log.Logger) PaymentController {
	return &paymentController{
		service: service,
		logger:  logger,
	}
}

func (c paymentController) CreatePayment(ctx *gin.Context) {
	var createPaymentRequest struct {
		CartID string `json:"cart_id"`
		Paid   bool   `json:"paid"`
	}
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	json.Unmarshal(data, &createPaymentRequest)

	// turns it into function
	claims := jwt.ExtractClaims(ctx)

	userID := claims["id"].(string)

	ok, err := c.service.CreatePayment(ctx, userID, createPaymentRequest.CartID)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, gin.H{"success": ok})
}

func (c paymentController) UpdatePayment(ctx *gin.Context) {

}
