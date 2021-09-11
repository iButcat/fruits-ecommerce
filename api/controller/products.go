package controller

import (
	"ecommerce/service"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

type ProductsController interface {
	GetById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type productsController struct {
	service service.ServiceProducts
	logger  log.Logger
}

func NewProductsController(service service.ServiceProducts, logger log.Logger) ProductsController {
	return &productsController{
		service: service,
		logger:  logger,
	}
}

func (c productsController) GetById(ctx *gin.Context) {

}

func (c productsController) GetAll(ctx *gin.Context) {
	products, err := c.service.GetProducts(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"products": products})
}

func (c productsController) Update(ctx *gin.Context) {
	var args = make(map[string]interface{})
	bodyDecoder := json.NewDecoder(ctx.Request.Body)
	if err := bodyDecoder.Decode(&args); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	log.Println("args: ", args)

}

func (c productsController) Delete(ctx *gin.Context) {

}
