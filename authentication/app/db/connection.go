package db

import (
	"github/madi-api/app/core"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func InitialConnection(dbName string, mongoURI string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		core.ErrorLogger(err.Error())
	}

	return client.Database(dbName)
}
