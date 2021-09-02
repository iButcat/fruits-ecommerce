package server

import (
	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(models.User)
			userGroup.POST("/register", user.RegisterUser)
		}
	}

	return router
}
