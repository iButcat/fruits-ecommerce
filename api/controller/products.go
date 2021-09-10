package controller

import (
	"ecommerce/service"
	"log"

	"github.com/gin-gonic/gin"
)

type ProductsController interface {
	GetById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
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

}

func (c productsController) Create(ctx *gin.Context) {

}

func (c productsController) Update(ctx *gin.Context) {

}

func (c productsController) Delete(ctx *gin.Context) {

}
