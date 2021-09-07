package router

import (
	"ecommerce/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func NewRouter(controller controllers.Controllers, logger log.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userGroup.POST("/register", controller.Register)
			userGroup.POST("/login", controller.Login)
		}

		productsGroup := v1.Group("products")
		{
			productsGroup.GET("/get", controller.GetProducts)
			productsGroup.PUT("/update", controller.UpdateProducts)
			productsGroup.POST("/create", controller.CreateProducts)
			productsGroup.DELETE("/delete", controller.DeleteProducts)
		}
	}

	return router
}
