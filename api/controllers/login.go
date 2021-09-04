package controllers

import (
	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

func (c controller) Login(ctx *gin.Context) {
	//username, password, ok := ctx.Request.BasicAuth()
	var userModel = models.User{}
	data, err := c.repository.Get(ctx, userModel, "1")
	if err != nil {
		ctx.JSON(400, gin.H{"error": err})
	}
	data = userModel
	//fmt.Println("is auth: ", ok)

	ctx.JSON(200, gin.H{"data": data})
}
