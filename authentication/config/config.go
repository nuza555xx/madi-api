package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost    string
	MongoUser     string
	MongoPassword string
	MongoHost     string
	MongoPort     string
	MongoDBName   string
	JwtSecret     string
	JwtExpiration string
}

func (config *Config) initialize() {
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.ServerHost = env["SERVER_HOST"]
	config.MongoUser = env["MONGO_USER"]
	config.MongoPassword = env["MONGO_PASSWORD"]
	config.MongoHost = env["MONGO_HOST"]
	config.MongoDBName = env["MONGO_DBNAME"]
	config.JwtSecret = env["JWT_SECRET"]
	config.JwtExpiration = env["JWT_EXPIRATION"]

}

func (config *Config) MongoURI() string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		config.MongoUser,
		config.MongoPassword,
		config.MongoHost,
	)
}
func (config *Config) MongoDB() string {
	return config.MongoDBName
}

func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
