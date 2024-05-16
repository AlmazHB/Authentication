package service

import (
	"github.com/diamondhulk625/web/internall/app/models"
	"github.com/diamondhulk625/web/internall/app/repository"
)

type Authorization interface {
	CreateUser(user models.CreateUserRequest) (string, error)
}
type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
