package router

import (
	"log"

	// internal pkg
	"ecommerce/controller"
	"ecommerce/middleware"
	"ecommerce/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type controllersRouter struct {
	authController     controller.AuthController
	productsController controller.ProductsController
	cartController     controller.CartController
	paymentController  controller.PaymentController
}

func NewControllerRouter(authController controller.AuthController,
	productsController controller.ProductsController,
	cartController controller.CartController,
	paymentController controller.PaymentController) *controllersRouter {
	return &controllersRouter{
		authController:     authController,
		productsController: productsController,
		cartController:     cartController,
		paymentController:  paymentController,
	}
}

func (cr controllersRouter) NewRouter(logger log.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	customJwtMiddleware := middleware.CustomJwtMiddleware(cr.authController)
	err := customJwtMiddleware.MiddlewareInit()
	if err != nil {
		log.Fatal("authMiddleware", err.Error())
	}

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userGroup.POST("/register", cr.authController.Register)
			userGroup.POST("/login", customJwtMiddleware.LoginHandler) // generate jwt token
		}

		authTestGroup := v1.Group("authtest")
		authTestGroup.Use(customJwtMiddleware.MiddlewareFunc())
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
		}

		cartGroup := v1.Group("cart")
		cartGroup.Use(customJwtMiddleware.MiddlewareFunc())
		{
			cartGroup.POST("/add", cr.cartController.Add)
			cartGroup.GET("/list", cr.cartController.List)
			cartGroup.PUT("/update", cr.cartController.Update)
		}

		paymentGroup := v1.Group("payment")
		paymentGroup.Use(customJwtMiddleware.MiddlewareFunc())
		{
			paymentGroup.POST("/payment", cr.paymentController.CreatePayment)
		}
	}
	return router
}
