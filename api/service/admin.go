package service

import (
	"context"
	"ecommerce/repository"
	"log"
)

type AdminService interface {
	GenerateCode(ctx context.Context) (bool, error)
	DeleteCode(ctx context.Context, id string) (bool, error)
}

type adminService struct {
	repository repository.Repository
	logger     log.Logger
}

func NewAdminService(repo repository.Repository, logger log.Logger) AdminService {
	return &adminService{
		repository: repo,
		logger:     logger,
	}
}

func (s adminService) GenerateCode(ctx context.Context) (bool, error) {
	return true, nil
}

func (s adminService) DeleteCode(ctx context.Context, id string) (bool, error) {
	return true, nil
}
