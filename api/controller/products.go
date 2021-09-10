package controller

import "github.com/gin-gonic/gin"

type ProductsController interface {
	GetById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
