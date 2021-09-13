package main

import (
	"ecommerce/models"
	"log"

	"gorm.io/gorm"
)

func initModels(db *gorm.DB, models ...interface{}) (bool, error) {
	err := db.Debug().AutoMigrate(models...)
	if err != nil {
		return false, err
	}

	for _, model := range models {
		tx := db.Debug().Create(model)
		if tx.Error != nil {
			return false, tx.Error
		}
	}

	return true, nil
}

func noot(db *gorm.DB) {
	ok, err := initModels(db, &models.Product{
		Name:     "bananas",
		Price:    0.85,
		Quantity: 99,
		Empty:    false,
	}, &models.Product{
		Name:     "apples",
		Price:    0.70,
		Quantity: 99,
		Empty:    false,
	}, &models.Product{
		Name:     "oranges",
		Price:    0.67,
		Quantity: 99,
		Empty:    false,
	}, &models.Product{
		Name:     "pears",
		Price:    0.85,
		Quantity: 99,
		Empty:    false,
	},
		&models.Products{},
		&models.Cart{})
	if ok && err == nil {
		log.Println("models has been created")
	} else {
		log.Println(err)
	}
}
