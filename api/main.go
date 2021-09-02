package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	// internal pkg
	"ecommerce/config"
	"ecommerce/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Println(err)
	}

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Println("error while initializing db...", err)
	}

	if db != nil {
		log.Println("db connected....")
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	go func() {
		log.Println("Starting server...")
		errs <- server.InitServer()
	}()

	log.Println("exit", <-errs)
}
