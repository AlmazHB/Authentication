package handler

import (
	"github.com/diamondhulk625/web/internall/app/service"
	"github.com/gin-gonic/gin"
)

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type ErrorResponse struct {
// 	Message string `json:"message"`
// }

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Если метод запроса OPTIONS, отправляем пустой ответ без выполнения обработки
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.GET("/sign-in", h.SignIn)
	}

	return r
}
