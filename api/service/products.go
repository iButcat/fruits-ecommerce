package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"log"
)

type ServiceProducts interface {
	GetProducts(ctx context.Context) (*models.Products, error)
	UpdateProducts(ctx context.Context)
	DeleteProducts(ctx context.Context)
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

func (s serviceProducts) GetProducts(ctx context.Context) (*models.Products, error) {
	data, err := s.repository.GetAll(ctx, &models.Products{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data.(*models.Products), nil
}

func (s serviceProducts) UpdateProducts(ctx context.Context) {

}

func (s serviceProducts) DeleteProducts(ctx context.Context) {

}
