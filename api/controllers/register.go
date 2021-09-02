package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	// internal pkg
	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

type User models.User

func (u User) RegisterUser(ctx *gin.Context) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("err while trying to register user: ", err)
	}
	if err := json.Unmarshal(data, &u); err != nil {
		log.Println("err while trying to unmarshal data: ", err)
	}
	fmt.Println(u)
}
