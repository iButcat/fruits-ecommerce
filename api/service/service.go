package service

import (
	"ecommerce/repository"
	"log"

	"github.com/gin-gonic/gin"
)

// Controllers for our different business logic
type Service interface {
	// auth controllers
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)

	// products controllers
	CreateProducts(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	UpdateProducts(ctx *gin.Context)
	DeleteProducts(ctx *gin.Context)
}

type service struct {
	repository repository.Repository
	logger     log.Logger
}

func NewService(repo repository.Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}
