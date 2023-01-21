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
	"golang.org/x/crypto/bcrypt"
)

func SignUpAccountWithEmail(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	account := new(model.Account)

	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Invalid request body")
		return
	}

	if ok, errors := core.ValidateInputs(account); !ok {
		ResponseWithValidationError(res, http.StatusBadRequest, errors)
		return
	}

	collection := db.Collection(core.AccountCollection)

	password, err := bcrypt.GenerateFromPassword([]byte(account.Password), 14)
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Error hashing password")
		return
	}
	account.Password = string(password)

	existingAccount := &model.Account{}
	err = collection.FindOne(context.TODO(), bson.M{"email": account.Email}).Decode(existingAccount)
	if err == nil {
		ResponseWithError(res, http.StatusBadRequest, "Email already exists")
		return
	}

	result, err := collection.InsertOne(context.TODO(), account)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			ResponseWithError(res, http.StatusBadRequest, "Email already exists")
		} else {
			ResponseWithError(res, http.StatusBadRequest, "Error inserting account into the database")
		}
		return
	}

	accessToken, err := core.GenerateJWT(map[string]interface{}{"_id": result.InsertedID.(primitive.ObjectID)})
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Error generating JWT")
		return
	}

	ResponseWithJSON(res, http.StatusCreated, map[string]interface{}{"accessToken": accessToken})
}

func SignInAccountWithEmail(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	signIn := new(model.SignInWithEmail)

	err := json.NewDecoder(req.Body).Decode(signIn)
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Invalid request body")
		return
	}

	if ok, errors := core.ValidateInputs(signIn); !ok {
		ResponseWithValidationError(res, http.StatusBadRequest, errors)
		return
	}

	collection := db.Collection(core.AccountCollection)

	var account model.Account
	err = collection.FindOne(context.TODO(), bson.M{"email": signIn.Email, "social": bson.M{"$exists": false}}).Decode(&account)
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Account not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(signIn.Password))
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Incorrect password")
		return
	}

	accessToken, err := core.GenerateJWT(map[string]interface{}{"_id": account.ID})
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Error generating JWT")
		return
	}

	ResponseWithJSON(res, http.StatusOK, map[string]interface{}{"accessToken": accessToken})
}

func SignInAccountWithSocial(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	account := new(model.AccountSyncSocial)

	err := json.NewDecoder(req.Body).Decode(account)
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Invalid request body")
		return
	}

	if ok, errors := core.ValidateInputs(account); !ok {
		ResponseWithValidationError(res, http.StatusBadRequest, errors)
		return
	}

	collection := db.Collection(core.AccountCollection)

	var existingAccount model.AccountSyncSocial
	err = collection.FindOne(context.TODO(), bson.M{"email": account.Email, "social.socialId": account.Social.SocialId}).Decode(&existingAccount)
	if err != nil {
		_, err = collection.InsertOne(context.TODO(), account)
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				ResponseWithError(res, http.StatusBadRequest, "Account already exists")
			} else {
				ResponseWithError(res, http.StatusBadRequest, "Error inserting account into the database")
			}
			return
		}
	}

	accessToken, err := core.GenerateJWT(map[string]interface{}{"_id": account.ID})
	if err != nil {
		ResponseWithError(res, http.StatusBadRequest, "Error generating JWT")
		return
	}

	ResponseWithJSON(res, http.StatusOK, map[string]interface{}{"accessToken": accessToken})
}
