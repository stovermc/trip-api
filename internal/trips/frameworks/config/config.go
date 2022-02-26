package config

import (
	"fmt"
	"os"
)

type Config struct {
	MongoDatabase string
	MongoHost     string
	MongoPort     string
	MongoUsername string
	MongoPassword string
}

func Init() *Config {
	mongoDatabase := requireEnv("MONGO_DATABASE")
	mongoHost := requireEnv("MONGO_HOST")
	mongoPort := requireEnv("MONGO_PORT")
	mongoUsername := requireEnv("MONGO_USERNAME")
	mongoPassword := requireEnv("MONGO_PASSWORD")

	fmt.Println(mongoDatabase)
	return &Config{
		MongoDatabase: mongoDatabase,
		MongoHost:     mongoHost,
		MongoPort:     mongoPort,
		MongoUsername: mongoUsername,
		MongoPassword: mongoPassword,
	}
}

func requireEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return ""
	}
	return value
}
