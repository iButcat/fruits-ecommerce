package router

import (
	"log"

	// internal pkg
	"ecommerce/controller"

	"github.com/gin-gonic/gin"
)

type controllersRouter struct {
	authController     controller.AuthController
	productsController controller.ProductsController
}

func NewControllerRouter(authController controller.AuthController,
	productsController controller.ProductsController) *controllersRouter {
	return &controllersRouter{
		authController:     authController,
		productsController: productsController,
	}
}

func (cr controllersRouter) NewRouter(logger log.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userGroup.POST("/register", cr.authController.Register)
			userGroup.POST("/login", cr.authController.Login)
		}

		productsGroup := v1.Group("products")
		{
			productsGroup.GET("/get", cr.productsController.GetById)
			productsGroup.PUT("/getall", cr.productsController.GetAll)
			productsGroup.POST("/create", cr.productsController.Update)
			productsGroup.DELETE("/delete", cr.productsController.Delete)
		}
	}

	return router
}
