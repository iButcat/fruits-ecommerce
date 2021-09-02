package main

import (
	"log"

	// internal pkg
	"ecommerce/config"

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
}
