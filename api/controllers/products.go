package controllers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func (c controller) GetProducts(ctx *gin.Context) {
	products, err := c.repository.GetAll(ctx, &models.Products{})
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"products": products})
}

func (c controller) CreateProducts(ctx *gin.Context) {
	products := new(models.Products)
	fmt.Println(products)
}

func (c controller) UpdateProducts(ctx *gin.Context) {
	var args = make(map[string]interface{})
	bodyDecoder := json.NewDecoder(ctx.Request.Body)
	if err := bodyDecoder.Decode(&args); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	log.Println("args: ", args)
}

func (c controller) DeleteProducts(ctx *gin.Context) {

}
