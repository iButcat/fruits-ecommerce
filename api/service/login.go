package service

import (
	"ecommerce/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	errWrongLoginCredentials = errors.New("err wrong credentials has been submitted")
)

func (s service) Login(ctx *gin.Context) {
	var userModel = models.User{}

	// TODO: move it into a function
	var args = make(map[string]interface{})

	readRequest, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	json.Unmarshal(readRequest, &userModel)

	json.Unmarshal(readRequest, &args)

	log.Println(args)
	data, err := s.repository.Get(ctx, &models.User{}, args)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// type convertion of our interface to our user models
	userRepo := data.(*models.User)

	if userRepo.Username != userModel.Username || userRepo.Password != userModel.Password {
		log.Println(errWrongLoginCredentials)
		ctx.JSON(400, gin.H{"error": errWrongLoginCredentials})
		return
	} else {

	}

	ctx.JSON(200, gin.H{"logged": true})
}
