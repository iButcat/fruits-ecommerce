package main

import (
	"gorm.io/gorm"
)

func createModels(db *gorm.DB, models ...interface{}) (bool, error) {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return false, err
		}

		tx2 := db.Model(model)
		if tx2.Error != nil {
			return false, tx2.Error
		}

		tx := db.Create(model)
		if tx.Error != nil {
			return false, tx.Error
		}
	}
	return true, nil
}
