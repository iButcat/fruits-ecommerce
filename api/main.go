package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	// internal pkg
	"ecommerce/config"
	"ecommerce/controllers"
	"ecommerce/repository"
	"ecommerce/router"

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

	var controller controllers.Controllers
	{
		repository := repository.NewRepo(db, log.Logger{})
		controller = controllers.NewControllers(repository, log.Logger{})
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	go func() {
		router := router.NewRouter(controller, log.Logger{})
		log.Println("Starting server...")
		errs <- router.Run()
	}()

	log.Println("exit", <-errs)
}
