package repository

import (
	"context"

	"github.com/AlmazHB/Authentication/internall/app/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthDB struct {
	db *mongo.Database
}

func NewAuthDB(db *mongo.Database) *AuthDB {
	return &AuthDB{db: db}
}

type Repository struct {
	AuthDB *AuthDB
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		AuthDB: NewAuthDB(db),
	}
}

func (a *AuthDB) InsertUser(user models.User) error {
	_, err := a.db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		logrus.Errorf("Error inserting user: %s", err.Error())
	}
	return err
}

func (a *AuthDB) FindUserByEmail(email string) (models.User, error) {
	var result models.User
	err := a.db.Collection("users").FindOne(context.Background(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		logrus.Errorf("Error finding user by email: %s", err.Error())
	}
	return result, err
}
