package handler

import (
	"context"
	"encoding/json"
	"github/madi-api/app/core"
	"github/madi-api/app/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccountWitEmail(db *mongo.Database, res http.ResponseWriter, req *http.Request) {

	account := new(model.Account)

	err := json.NewDecoder(req.Body).Decode(account)
	if err != nil {
		ResponseWriter(res, http.StatusBadRequest, "Body json request have issues!!!", nil)
		return
	}

	if ok, errors := core.ValidateInputs(account); !ok {
		model.NewValidatedResponse(http.StatusBadRequest, "Invalid is value please check again.", errors)
		return
	}

	collection := db.Collection("account")

	password, err := bcrypt.GenerateFromPassword([]byte(account.Password), 14)
	account.Password = string(password)

	if err != nil {
		ResponseWriter(res, http.StatusBadRequest, "Password hash is failed.", nil)
		return
	}

	_, err = collection.InsertOne(context.TODO(), model.NewAccount(account))
	if err != nil {
		mongoException := err.(mongo.WriteException)
		if mongoException.WriteErrors[0].Code == 11000 {
			ResponseWriter(res, http.StatusBadRequest, "Email is existing, Please try again.", nil)
		} else {
			ResponseWriter(res, http.StatusBadRequest, err.Error(), nil)
		}
		return
	}

	findOptions := options.FindOne().SetProjection(bson.D{primitive.E{Key: "email", Value: 1}})
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: account.Email}}, findOptions).Decode(&account)
	if err != nil {
		ResponseWriter(res, http.StatusBadRequest, "Account does not exist.", nil)
		return
	}

	accessToken, err := core.GenerateJWT(map[string]interface{}{"_id": account.ID})

	if err != nil {
		ResponseWriter(res, http.StatusBadRequest, "Invalid authentication token.", nil)
		return
	}

	ResponseWriter(res, http.StatusCreated, "", map[string]interface{}{"accessToken": accessToken})

}
