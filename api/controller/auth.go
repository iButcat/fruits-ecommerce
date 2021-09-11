package controller

import (
	"ecommerce/models"
	"ecommerce/service"
	"ecommerce/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context) (interface{}, error)
	Logout(ctx *gin.Context)
}

type authController struct {
	service service.AuthService
	logger  log.Logger
}

func NewAuthController(service service.AuthService, logger log.Logger) AuthController {
	return &authController{
		service: service,
		logger:  logger,
	}
}

var (
	errWrongLoginCredentials = errors.New("err wrong credentials has been submitted")
	errTokenIsNotValid       = errors.New("err token is not valid")
)

func (c authController) Register(ctx *gin.Context) {
	var user = new(models.User)
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("err while trying to read request: ", err)
		return
	}

	if err := json.Unmarshal(data, &user); err != nil {
		log.Println("err while trying to unmarshal data: ", err)
		return
	}

	var lengthPassword = len(user.Password)
	var lengthUsername = len(user.Username)
	if !utils.ValidateIfNotEmpty(lengthPassword, lengthUsername) {
		ctx.JSON(400, gin.H{"error": "missing fields password or username"})
		ctx.Abort()
		return
	}

	if !utils.RegexEmailChecker(user.Email) {
		ctx.JSON(400, gin.H{"error": "email is not valid"})
		ctx.Abort()
		return
	}

	saveUser, err := c.service.Register(ctx, *user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort()
		log.Println("error while saving user: ", err)
		return
	}

	ctx.JSON(201, gin.H{"success": saveUser})
}

// modify return value, not that great
func (c authController) Login(ctx *gin.Context) (interface{}, error) {
	var userModel = models.User{}

	// TODO: move it into a function
	var args = make(map[string]interface{})

	readRequest, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return nil, err
	}

	json.Unmarshal(readRequest, &userModel)

	json.Unmarshal(readRequest, &args)

	log.Println(args)
	userRepo, err := c.service.Login(ctx, userModel.Username, userModel.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return nil, err
	}

	if userRepo.Username != userModel.Username || userRepo.Password != userModel.Password {
		log.Println(errWrongLoginCredentials)
		ctx.JSON(400, gin.H{"error": errWrongLoginCredentials})
		return nil, err
	}

	ctx.JSON(200, gin.H{"logged": true})
	return &userRepo, nil
}

func (c authController) Logout(ctx *gin.Context) {

}
