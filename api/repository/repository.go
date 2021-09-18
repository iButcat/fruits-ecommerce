package repository

import (
	"context"
	"log"

	// test

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Generic repo for our differents models
type Repository interface {
	Create(ctx context.Context, models interface{}) (string, error)
	GetRows(ctx context.Context, models interface{}) (interface{}, error)
	Get(ctx context.Context, models interface{}, fields map[string]interface{}) (interface{}, error)
	GetAll(ctx context.Context, models interface{}) (interface{}, error)
	Update(ctx context.Context, models interface{}, id string, fields map[string]interface{}) (bool, error)
	UpdateNested(ctx context.Context, models interface{}) (bool, error)
	Delete(ctx context.Context, models interface{}, id string) (bool, error)
}

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

// Create data from any given models
func (repo *repo) Create(ctx context.Context, models interface{}) (string, error) {
	if err := repo.db.Model(models).Error; err != nil {
		return "", err
	}

	tx := repo.db.Create(models)
	if tx.Error != nil {
		return "err while creating models: ", tx.Error
	}
	return "Data has been created", nil
}

func (repo *repo) GetRows(ctx context.Context, models interface{}) (interface{}, error) {
	rows, err := repo.db.Model(models).Preload(clause.Associations).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		repo.db.ScanRows(rows, models)
	}
	return models, nil
}

// Get data from different fields like id, username, password, etc...
func (repo *repo) Get(
	ctx context.Context, models interface{}, fields map[string]interface{}) (interface{}, error) {
	log.Println("FIELDS REPO: ", fields)
	if err := repo.db.Where(fields).Preload(clause.Associations).Find(models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

// Get All
func (repo *repo) GetAll(ctx context.Context, models interface{}) (interface{}, error) {
	if err := repo.db.Preload(clause.Associations).Find(models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

// update any given models with their column and values that need to be change
func (repo *repo) Update(
	ctx context.Context, models interface{}, id string, fields map[string]interface{}) (bool, error) {
	for index, value := range fields {
		if err := repo.db.Model(models).Where("id = ?", id).Update(index, value).Error; err != nil {
			return false, err
		}
	}
	return true, nil
}

func (repo *repo) UpdateNested(ctx context.Context, models interface{}) (bool, error) {
	if err := repo.db.Debug().Save(models).Error; err != nil {
		return false, err
	}
	return true, nil
}

// delete any given data from models with id
func (repo *repo) Delete(ctx context.Context, models interface{}, id string) (bool, error) {
	if err := repo.db.Delete(models, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
