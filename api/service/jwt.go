package service

import (
	"ecommerce/models"
	"ecommerce/repository"
	"log"
)

type JwtService interface {
	GenerateToken(user models.User, isAdmin bool) (models.Jwt, error)
	ValidateToken(token models.Jwt) (bool, error)
}

type jwtService struct {
	logger     log.Logger
	repository repository.Repository
}

func NewJwtService(repo repository.Repository, logger log.Logger) JwtService {
	return &jwtService{
		logger:     logger,
		repository: repo,
	}
}

func (s jwtService) GenerateToken(user models.User, isAdmin bool) (models.Jwt, error) {
	return models.Jwt{}, nil
}

func (s jwtService) ValidateToken(token models.Jwt) (bool, error) {
	return true, nil
}
