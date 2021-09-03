package repository

import (
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type repo struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

/*
// Create table for our models if not exists
func (repo *repo) createMigrationModel(model interface{}) {
	fmt.Println("create migration model value: ", model)
	tx := repo.db.AutoMigrate(model)
	if tx.Error != nil {
		repo.logger.Println("err while creating table: ", tx.Error)
		return
	}
}
*/

func (repo *repo) Create(ctx context.Context, models interface{}) (string, error) {
	fmt.Println("create repo models value: ", models)
	repo.db.AutoMigrate(models)
	tx2 := repo.db.Model(models)
	if tx2.Error != nil {
		return "err while creating table: ", tx2.Error
	}

	tx := repo.db.Create(models)
	if tx.Error != nil {
		return "err while creating models: ", tx.Error
	}
	return "Data has been created", nil
}

func (repo *repo) Get(ctx context.Context, fields ...string) (interface{}, error) {
	return nil, nil
}

func (repo *repo) GetAll(ctx context.Context) ([]interface{}, error) {
	return nil, nil
}

func (repo *repo) Update(ctx context.Context, models interface{}) (bool, error) {
	return true, nil
}

func (repo *repo) Delete(ctx context.Context, id string) (bool, error) {
	return true, nil
}
