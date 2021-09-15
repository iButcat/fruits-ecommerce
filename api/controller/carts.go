package controller

import (
	"io/ioutil"
	"log"

	// internal pkg
	"ecommerce/models"
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
	service service.ServiceCarts
	logger  log.Logger
}

func NewCartController(service service.ServiceCarts, logger log.Logger) CartController {
	return &cartController{
		service: service,
		logger:  logger,
	}
}

func (c cartController) Add(ctx *gin.Context) {
	cart := models.Cart{}
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	json.Unmarshal(data, &cart)
	claims := jwt.ExtractClaims(ctx)
	var fields = make(map[string]string)
	fields["username"] = claims["id"].(string)
	fields["product_name"] = "bananas"

	success, err := c.service.AddCarts(ctx, cart, fields["product_name"], fields["username"])
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

}
