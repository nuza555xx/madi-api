package auth

import (
	"context"
	"encoding/json"
	"madi-api/common"
	adapter "madi-api/config"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db = adapter.Connect().Database(common.DotEnvVariable("DB_NAME"))

func registerWithEmail() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		var register RegisterWithEmailDto

		err := json.NewDecoder(request.Body).Decode(&register)
		if err != nil {
			common.ErrorResponse(err.Error(), response)
			return
		}

		if ok, errors := common.ValidateInputs(register); !ok {
			common.ValidationResponse(errors, response)
			return
		}

		collection := db.Collection("account")

		// _, errIndex := collection.Indexes().CreateOne(
		// 	context.Background(),
		// 	mongo.IndexModel{
		// 		Keys: bson.M{
		// 			"email": 1,
		// 		},
		// 		Options: options.Index().SetUnique(true),
		// 	},
		// )
		// if err != nil {
		// 	common.ErrorLogger(errIndex)
		// 	return
		// }

		password, err := common.HashPassword(register.Password)

		if err != nil {
			common.ErrorResponse(err.Error(), response)
			return
		}

		docs := bson.D{
			primitive.E{Key: "email", Value: register.Email},
			primitive.E{Key: "password", Value: password},
			primitive.E{Key: "phone", Value: register.Phone},
			primitive.E{Key: "displayName", Value: register.DisplayName},
			primitive.E{Key: "createdAt", Value: time.Now()},
			primitive.E{Key: "updatedAt", Value: time.Now()},
		}

		_, err = collection.InsertOne(context.Background(), docs)
		if err != nil {
			mongoException := err.(mongo.WriteException)
			if mongoException.WriteErrors[0].Code == 11000 {
				common.ErrorResponse("Email is existing, Please try again.", response)
			} else {
				common.ErrorResponse(err.Error(), response)
			}
			return
		}

		// var account Account

		// findOptions := options.FindOne().SetProjection(bson.D{primitive.E{Key: "email", Value: 1}})

		// err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: register.Email}}, findOptions).Decode(&account)
		// if err != nil {
		// 	common.ErrorResponse("Account does not exist.", response)
		// 	return
		// }

		// accessToken, err := common.GenerateJWT(account)

		// if err != nil {
		// 	common.ErrorResponse("Invalid authentication token.", response)
		// 	return
		// }

		// common.SuccessResponse(map[string]interface{}{"accessToken": accessToken}, response)
	})
}
