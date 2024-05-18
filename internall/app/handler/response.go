package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	fmt.Println(message)
	ctx.AbortWithStatusJSON(statusCode, error{message})
}
