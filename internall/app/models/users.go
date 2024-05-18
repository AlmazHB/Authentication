package models

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	County   string `json:"county" bson:"county"`
	Name     string `json:"name" bson:"name"`
}

type CreateUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	County   string `json:"county"`
	Name     string `json:"name"`
}
