package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost    string // address that server will listening on
	MongoUser     string // mongo db username
	MongoPassword string // mongo db password
	MongoHost     string // host that mongo db listening on
	MongoPort     string // port that mongo db listening on
	JwtSecret     string // jwt secret

}

func (config *Config) initialize() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.ServerHost = os.Getenv("SERVER_HOST")
	config.MongoUser = os.Getenv("MONGO_USER")
	config.MongoPassword = os.Getenv("MONGO_PASSWORD")
	config.MongoHost = os.Getenv("MONGO_HOST")
	config.MongoPort = os.Getenv("MONGO_PORT")
	config.JwtSecret = os.Getenv("JWT_SECRET")

}

func (config *Config) MongoURI() string {
	return fmt.Sprintf("mongodb://%s:%s",
		config.MongoHost,
		config.MongoPort,
	)
}

func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
