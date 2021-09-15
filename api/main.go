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

	var (
		repository         = repository.NewRepo(db, log.Logger{})
		authService        = service.NewAuthService(repository, log.Logger{})
		authController     = controller.NewAuthController(authService, log.Logger{})
		productsService    = service.NewServiceProducts(repository, log.Logger{})
		productsController = controller.NewProductsController(productsService, log.Logger{})
		cartService        = service.NewServiceCarts(repository, log.Logger{})
		cartController     = controller.NewCartController(cartService, log.Logger{})
		routerController   = router.NewControllerRouter(authController, productsController, cartController)
	)

	//noot(db)
	//noot1(db)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	go func() {
		router := routerController.NewRouter(log.Logger{})
		log.Println("Starting server...")
		errs <- router.Run(config.Port)
		log.Println(errs)
	}()

	log.Println("exit", <-errs)
}
