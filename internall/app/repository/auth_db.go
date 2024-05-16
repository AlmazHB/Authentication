package repository

import (
	"net/http"

	"github.com/diamondhulk625/web/internall/app/database"
	"github.com/diamondhulk625/web/internall/app/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthDB struct {
	db *mongo.Database
}

func NewAuthDB(db *mongo.Database) *AuthDB {
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(input models.CreateUserRequest) (string, error) {
	var ctx *gin.Context
	res, err := database.Users.InsertOne(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add user"})
	}
	user := models.User{
		ID:       res.InsertedID.(primitive.ObjectID),
		Name:     input.Name,
		LastName: input.LastName,
		Country:  input.Country,
		Email:    input.Email,
		Password: input.Password,
	}
	return user.ID.String(), nil
}
