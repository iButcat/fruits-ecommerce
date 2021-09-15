package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"log"
)

type ServiceProducts interface {
	GetProducts(ctx context.Context) (*[]models.Product, error)
	GetProduct(ctx context.Context, id string) (*models.Product, error)
}

type serviceProducts struct {
	repository repository.Repository
	logger     log.Logger
}

func NewServiceProducts(repo repository.Repository, logger log.Logger) ServiceProducts {
	return &serviceProducts{
		repository: repo,
		logger:     logger,
	}
}

func (s serviceProducts) GetProducts(ctx context.Context) (*[]models.Product, error) {
	allProducts := []models.Product{}
	data, err := s.repository.GetRows(ctx, &allProducts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data.(*[]models.Product), nil
}

func (s serviceProducts) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	return nil, nil
}
