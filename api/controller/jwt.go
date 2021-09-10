package controller

import "github.com/gin-gonic/gin"

type JwtController interface {
	GenerateToken(ctx *gin.Context)
	ValidateToken(ctx *gin.Context)
}
