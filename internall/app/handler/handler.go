package handler

import (
	"errors" // Добавляем этот импорт для работы с errors.New
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/AlmazHB/Authentication/internall/app/models"
	"github.com/AlmazHB/Authentication/internall/app/service"
)

type Handler struct {
	authService *service.AuthService
}

func NewHandler(authService *service.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/v1/api")
	{
		api.POST("/register", h.register)
		api.GET("/login", h.login)
		api.POST("/reset", h.resetPassword)
		api.GET("/welcome", h.authenticate(h.welcome))
	}

	return router
}

func (h *Handler) register(c *gin.Context) {
	var creds models.CreateUserRequest
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.authService.RegisterUser(creds); err != nil {
		logrus.Errorf("Error registering user: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *Handler) login(c *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.authService.AuthenticateUser(creds.Email, creds.Password)
	if err != nil {
		logrus.Errorf("Error authenticating user: %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) resetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Password reset endpoint"})
}

func getJWTSecret() ([]byte, error) {
	secret := viper.GetString("jwtSecret")
	if secret == "" {
		return nil, errors.New("JWT secret is not configured")
	}
	return []byte(secret), nil
}

func (h *Handler) authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		claims := &service.Claims{}
		secret, err := getJWTSecret()
		if err != nil {
			logrus.Errorf("Error getting JWT secret: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			logrus.Errorf("Error validating token: %s", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		next(c)
	}
}

func (h *Handler) welcome(c *gin.Context) {
	email := c.MustGet("email").(string)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + email})
}
