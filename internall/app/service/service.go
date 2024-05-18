package service

import (
	"time"

	"github.com/AlmazHB/Authentication/internall/app/models"
	"github.com/AlmazHB/Authentication/internall/app/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthDB
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo.AuthDB,
	}
}

func (s *AuthService) RegisterUser(creds models.CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    creds.Email,
		Password: string(hashedPassword),
		County:   creds.County,
		Name:     creds.Name,
	}
	return s.repo.InsertUser(user)
}

func (s *AuthService) AuthenticateUser(email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	storedPassword := user.Password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("jwtSecret")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
