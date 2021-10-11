package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	// internal pkg
	"ecommerce/config"
	"ecommerce/controller"
	"ecommerce/repository"
	"ecommerce/router"
	"ecommerce/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Println(err)
	}

	var db *gorm.DB
	{
		var err error
		db, err = gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
		if err != nil {
			log.Println("error while initializing db...", err)
		}
	}

	var logger = log.Logger{}

	var (
		repository         = repository.NewRepo(db, logger)
		authService        = service.NewAuthService(repository, logger)
		authController     = controller.NewAuthController(authService, logger)
		productsService    = service.NewServiceProducts(repository, logger)
		productsController = controller.NewProductsController(productsService, logger)
		cartService        = service.NewServiceCarts(repository, logger)
		cartController     = controller.NewCartController(cartService, logger)
		paymentService     = service.NewServicePayment(repository, logger)
		paymentController  = controller.NewPaymentController(paymentService, logger)
		routerController   = router.NewControllerRouter(authController, productsController, cartController, paymentController)
	)

	//noot(db)
	//noot1(db)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	go func() {
		router := routerController.NewRouter(logger)
		log.Println("Starting server...")
		errs <- router.Run(config.Port)
		log.Println(errs)
	}()

	log.Println("exit", <-errs)
}
