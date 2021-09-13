package controller

import (
	"fmt"
	"io/ioutil"
	"log"

	// internal pkg
	"ecommerce/models"
	"ecommerce/service"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type CartController interface {
	Add(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

	user, _ := ctx.Get("id")
	userId := user.(*models.User).ID

	success, err := c.service.AddCarts(ctx, cart, fmt.Sprint(userId))
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": success})
}

func (c cartController) List(ctx *gin.Context) {

}

func (c cartController) Update(ctx *gin.Context) {

}

func (c cartController) Delete(ctx *gin.Context) {

}
