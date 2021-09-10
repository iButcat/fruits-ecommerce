package router

import (
	"log"

	// internal pkg
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(service service.Service, logger log.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userGroup.POST("/register", service.Register)
			userGroup.POST("/login", service.Login)
		}

		productsGroup := v1.Group("products")
		{
			productsGroup.GET("/get", service.GetProducts)
			productsGroup.PUT("/update", service.UpdateProducts)
			productsGroup.POST("/create", service.CreateProducts)
			productsGroup.DELETE("/delete", service.DeleteProducts)
		}
	}

	return router
}
