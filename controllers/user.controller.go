package controllers

import (
	"context"
	"go-mongo-api/configs"
	"go-mongo-api/models"
	"go-mongo-api/responses"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

var validate = validator.New()

func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	// validate the request body

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Status: http.StatusBadRequest, Message: "error create", Data: &echo.Map{"data": err.Error()}})
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}
	password, err := models.HashPassword(user.Password)

	if err != nil {
		return err
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Username: user.Username,
		Password: password,
		Email:    user.Email,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: &echo.Map{"data": err.Error()}})
	}
	var info models.User
	err = userCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&info)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": info}})
}

func Login(c echo.Context) error {

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Status: http.StatusAccepted, Message: "success", Data: &echo.Map{"data": "info"}})
}
