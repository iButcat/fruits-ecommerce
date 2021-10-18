package controller

import (
	"ecommerce/service"
	"log"

	"github.com/gin-gonic/gin"
)

type AdminController interface {
	GenerateCode(ctx *gin.Context)
	DeleteCode(ctx *gin.Context)
}

type adminController struct {
	service service.AdminService
	logger  log.Logger
}

func NewAdminController(service service.AdminService, logger log.Logger) AdminController {
	return &adminController{
		service: service,
		logger:  logger,
	}
}

func (c adminController) GenerateCode(ctx *gin.Context) {
	//var generateCodeRequest struct{}
}

func (c adminController) DeleteCode(ctx *gin.Context) {
	//var deleteCodeRequest struct{}
}
