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
	}

	return router
}
