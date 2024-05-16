package repository

import (
	"github.com/diamondhulk625/web/internall/app/models"
)

type Authorization interface {
	CreateUser(user models.CreateUserRequest) (string, error)
}
type Repository struct {
	Authorization
}

func NewRepository() *Repository {
	return &Repository{
		Authorization: NewAuthDB(),
	}
}
