package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name" `
	LastName string             `bson:"lastname" json:"lastname" `
	Country  string             `bson:"country" json:"country" `
	Email    string             `bson:"email" json:"email" `
	Password string             `bson:"password" json:"pasword" `
}

type CreateUserRequest struct {
	Name     string `bson:"name" json:"name"`
	LastName string `bson:"lastname" json:"lastname"`
	Country  string `bson:"country" json:"country"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"pasword"`
}
