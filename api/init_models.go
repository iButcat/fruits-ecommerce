package main

import (
	"ecommerce/models"
	"log"

	"gorm.io/gorm"
)

func migrateModels(db *gorm.DB, models ...interface{}) (bool, error) {
	for _, model := range models {
		if err := db.Debug().AutoMigrate(model); err != nil {
			return false, err
		}
	}
	return true, nil
}

func createBaseData(db *gorm.DB, models ...interface{}) (bool, error) {
	for _, model := range models {
		if err := db.Debug().Create(model).Error; err != nil {
			return false, err
		}
	}
	return true, nil
}

func noot(db *gorm.DB) {
	ok, err := migrateModels(db,
		&models.Product{},
		&models.User{},
		&models.Cart{},
		&models.CartItem{},
		&models.Payment{})
	if ok && err == nil {
		log.Println("models has been created")
	} else {
		log.Println(err)
	}
}

func noot1(db *gorm.DB) {
	ok, err := createBaseData(db, &models.Product{
		Name:  "bananas",
		Price: 0.85,
		Empty: false,
	}, &models.Product{
		Name:  "apples",
		Price: 0.70,
		Empty: false,
	}, &models.Product{
		Name:  "oranges",
		Price: 0.67,
		Empty: false,
	}, &models.Product{
		Name:  "pears",
		Price: 0.85,
		Empty: false,
	})
	if ok && err == nil {
		log.Println("models has been created")
	} else {
		log.Println(err)
	}
}
