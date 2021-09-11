package router

import (
	"log"

	// internal pkg
	"ecommerce/controller"
	"ecommerce/middleware"
	"ecommerce/models"

	jwt "github.com/appleboy/gin-jwt/v2"
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

	err := middleware.CustomJwtMiddleware().MiddlewareInit()
	if err != nil {
		log.Fatal("authMiddleware", err.Error())
	}

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userGroup.POST("/register", cr.authController.Register)
			userGroup.POST("/login", cr.authController.Login)
			userGroup.POST("/token", middleware.CustomJwtMiddleware().LoginHandler) // generate jwt token
		}

		authTestGroup := v1.Group("authtest")
		authTestGroup.Use(middleware.CustomJwtMiddleware().MiddlewareFunc())
		{
			authTestGroup.GET("/hello", func(ctx *gin.Context) {
				claims := jwt.ExtractClaims(ctx)
				user, _ := ctx.Get("id")
				ctx.JSON(200, gin.H{"userId": claims["id"],
					"username": user.(*models.User).Username,
					"text":     "Hello World"})
			})
		}

		productsGroup := v1.Group("products")
		{
			productsGroup.GET("/get", cr.productsController.GetById)
			productsGroup.GET("/getall", cr.productsController.GetAll)
			productsGroup.PATCH("/update", cr.productsController.Update)
			productsGroup.DELETE("/delete", cr.productsController.Delete)
		}
	}
	return router
}
