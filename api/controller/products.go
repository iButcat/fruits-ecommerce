package controller

import (
	"ecommerce/service"
	"log"

	"github.com/gin-gonic/gin"
)

type ProductsController interface {
	GetById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type productsController struct {
	service service.ProductsService
	logger  log.Logger
}

func NewProductsController(service service.ProductsService, logger log.Logger) ProductsController {
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
