package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	// internal pkg
	"ecommerce/models"
	"ecommerce/utils.go"

	"github.com/gin-gonic/gin"
)

func (c controller) Register(ctx *gin.Context) {
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

	ok, err := c.repository.Create(ctx, user)
	if err != nil {
		log.Println("error while saving user: ", err)
		return
	}
	fmt.Println(ok)

	ctx.JSON(201, gin.H{"message": "User has been created"})
}
