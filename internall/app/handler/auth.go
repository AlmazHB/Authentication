package handler

import (
	"net/http"

	"github.com/diamondhulk625/web/internall/app/database"
	"github.com/diamondhulk625/web/internall/app/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) SignUp(ctx *gin.Context) {

	var body models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	// id, err := h.services.Authorization.CreateUser(body)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add user"})
	// }
	res, err := database.Users.InsertOne(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add user"})
	}
	user := models.User{
		ID:       res.InsertedID.(primitive.ObjectID),
		Name:     body.Name,
		LastName: body.LastName,
		Country:  body.Country,
		Email:    body.Email,
		Password: body.Password,
	}
	ctx.JSON(http.StatusOK, user.ID)
}
func (h *Handler) SignIn(ctx *gin.Context) {

	cursor, err := database.Users.Find(ctx, bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch user"})
		return
	}

	var user []models.User
	if err = cursor.All(ctx, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
