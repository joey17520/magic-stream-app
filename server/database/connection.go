package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: unable to found .env file")
	}

	MongoDB := os.Getenv("MONGODB_URI")
	if MongoDB == "" {
		log.Fatal("MONGODB_URI not set!")
	}

	fmt.Println("MongoDB URI: ", MongoDB)

	clientOptions := options.Client().ApplyURI(MongoDB)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Failed to connect Mongodb: ", err)
	}

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: unable to found .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	collection := Client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		return nil
	}

	return collection
}
