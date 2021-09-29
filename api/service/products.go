package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"log"
)

type ProductsService interface {
	GetProducts(ctx context.Context) (*[]models.Product, error)
	GetProduct(ctx context.Context, id string) (*models.Product, error)
}

type productsService struct {
	repository repository.Repository
	logger     log.Logger
}

func NewServiceProducts(repo repository.Repository, logger log.Logger) ProductsService {
	return &productsService{
		repository: repo,
		logger:     logger,
	}
}

func (s productsService) GetProducts(ctx context.Context) (*[]models.Product, error) {
	allProducts := []models.Product{}
	data, err := s.repository.GetRows(ctx, &allProducts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data.(*[]models.Product), nil
}

func (s productsService) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	return nil, nil
}
