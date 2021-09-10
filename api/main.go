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

	repository := repository.NewRepo(db, log.Logger{})
	// clean up with var ()
	authService := service.NewAuthService(repository, log.Logger{})
	authController := controller.NewAuthController(authService, log.Logger{})
	productsService := service.NewServiceProducts(repository, log.Logger{})
	productsController := controller.NewProductsController(productsService, log.Logger{})

	var routerController = router.NewControllerRouter(authController, productsController)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	go func() {
		router := routerController.NewRouter(log.Logger{})
		log.Println("Starting server...")
		errs <- router.Run()
		log.Println(errs)
	}()

	log.Println("exit", <-errs)
}
