package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	// internal pkg
	"ecommerce/models"
	"ecommerce/utils"

	"github.com/gin-gonic/gin"
)

func (s service) Register(ctx *gin.Context) {
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

	ok, err := s.repository.Create(ctx, user)
	if err != nil {
		log.Println("error while saving user: ", err)
		return
	}
	fmt.Println(ok)

	ctx.JSON(201, gin.H{"message": "User has been created"})
}
