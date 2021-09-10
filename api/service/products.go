package service

import (
	"ecommerce/models"
	"ecommerce/repository"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type ServiceProducts interface {
	CreateProducts(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	UpdateProducts(ctx *gin.Context)
	DeleteProducts(ctx *gin.Context)
}

type serviceProducts struct {
	repository repository.Repository
	logger     log.Logger
}

func NewServiceProducts(repo repository.Repository, logger log.Logger) ServiceProducts {
	return &serviceProducts{
		repository: repo,
		logger:     logger,
	}
}

func (s serviceProducts) GetProducts(ctx *gin.Context) {
	products, err := s.repository.GetAll(ctx, &models.Products{})
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"products": products})
}

func (s serviceProducts) CreateProducts(ctx *gin.Context) {
	products := new(models.Products)
	fmt.Println(products)
}

func (s serviceProducts) UpdateProducts(ctx *gin.Context) {
	var args = make(map[string]interface{})
	bodyDecoder := json.NewDecoder(ctx.Request.Body)
	if err := bodyDecoder.Decode(&args); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	log.Println("args: ", args)
}

func (s serviceProducts) DeleteProducts(ctx *gin.Context) {

}
