package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/diamondhulk625/web/internall/app/models"
	"github.com/diamondhulk625/web/internall/app/repository"
)

const salt = "gdbcncdklsmcl457ndv"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user models.CreateUserRequest) (string, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
