package service

import (
	"context"
	"ecommerce/models"
	"ecommerce/repository"
	"ecommerce/utils"
	"errors"
	"log"
)

type AuthService interface {
	Register(ctx context.Context, user models.User) (string, error)
	Login(ctx context.Context, username, password string) (models.User, error)
	Logout(ctx context.Context) (bool, error)
}

type authService struct {
	repository repository.Repository
	logger     log.Logger
}

func NewAuthService(repo repository.Repository, logger log.Logger) AuthService {
	return &authService{
		repository: repo,
		logger:     logger,
	}
}

var (
	errWrongEmail  = errors.New("error email submited is not correct")
	errWhileSaving = errors.New("err while saving user")
)

// register service call repo to save our user
func (s authService) Register(ctx context.Context, user models.User) (string, error) {
	var lengthPassword = len(user.Password)
	var lengthUsername = len(user.Username)
	if !utils.ValidateIfNotEmpty(lengthPassword, lengthUsername) {
		return "", errWrongEmail
	}

	if !utils.RegexEmailChecker(user.Email) {
		return "", errWhileSaving
	}

	var fielUsername = make(map[string]interface{})
	fielUsername["username"] = user.Username
	dataUser, err := s.repository.Get(ctx, &models.User{}, fielUsername)
	if err != nil {
		return "error while fetching user data: ", err
	}
	checkUsername := dataUser.(*models.User).Username

	var errUsernameExists = errors.New("err username is already taken")
	if checkUsername == user.Username {
		return "", errUsernameExists
	}

	var fieldEmail = make(map[string]interface{})
	fieldEmail["email"] = user.Email
	dataUserEmail, err := s.repository.Get(ctx, &models.User{}, fieldEmail)
	if err != nil {
		return "error while fetching user data: ", err
	}
	checkEmail := dataUserEmail.(*models.User).Email
	var errEmailExists = errors.New("err email is already taken")
	if checkEmail == user.Email {
		return "", errEmailExists
	}

	ok, err := s.repository.Create(ctx, &user)
	if err != nil {
		log.Println("error while saving user: ", err)
		return "error while saving user: ", err
	}
	return ok, nil
}

func (s authService) Login(ctx context.Context, username, password string) (models.User, error) {
	var column = make(map[string]interface{})
	column["username"] = &username
	column["password"] = &password
	data, err := s.repository.Get(context.Background(), &models.User{}, column)
	if err != nil {
		return models.User{}, nil
	}

	// type convertion of our interface to our user models
	userRepo := data.(*models.User)

	return *userRepo, nil
}

func (s authService) Logout(ctx context.Context) (bool, error) {
	return true, nil
}
