package controllers

import (
	"ecommerce/repository"
	"log"

	"github.com/gin-gonic/gin"
)

// Controllers for our different business logic
type Controllers interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type controller struct {
	repository repository.Repository
	logger     log.Logger
}

func NewControllers(repo repository.Repository, logger log.Logger) Controllers {
	return &controller{
		repository: repo,
		logger:     logger,
	}
}
