package repository

import (
	"context"
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

// Create table for our models if not exists
func createMigrationModel(repo *repo, model interface{}) error {
	err := repo.db.AutoMigrate(model)
	if err != nil {
		return err
	}
	return nil
}

// Create data from any given models
func (repo *repo) Create(ctx context.Context, models interface{}) (string, error) {
	// auto migrate if need it.
	if err := createMigrationModel(repo, models); err != nil {
		return "", err
	}

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

// Get data from different fields like id, username, password, etc...
func (repo *repo) Get(ctx context.Context, models interface{}, fields ...string) (interface{}, error) {
	var args []string
	var id string
	var username string
	var password string

	args = append(args, fields...)

	if len(args) != 1 {
		username = args[0]
		password = args[1]
		if err := repo.db.Where("username = ? AND password = ? ",
			username, password).Find(models).Error; err != nil {
			log.Println("error while querying: ", err)
		}
	} else {
		id = args[0]
		if err := repo.db.First(models, id).Error; err != nil {
			log.Println("error while querying: ", err)
		}
	}
	return models, nil
}

func (repo *repo) GetAll(ctx context.Context, models interface{}) (interface{}, error) {
	if err := repo.db.Find(models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (repo *repo) Update(ctx context.Context, models interface{}, fields map[string]interface{}) (bool, error) {
	for index, value := range fields {
		if err := repo.db.Model(models).Update(index, value).Error; err != nil {
			return false, err
		}
	}
	return true, nil
}

func (repo *repo) Delete(ctx context.Context, models interface{}, id string) (bool, error) {
	if err := repo.db.Delete(models, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
