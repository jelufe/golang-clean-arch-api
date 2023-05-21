package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {
	envError := godotenv.Load(".env")

	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	mongoDb := os.Getenv("MONGODB_URL")

	client, mongoError := mongo.NewClient(options.Client().ApplyURI(mongoDb))

	if mongoError != nil {
		log.Fatal(mongoError)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	connectError := client.Connect(ctx)

	if connectError != nil {
		log.Fatal(connectError)
	}

	fmt.Println("Connected to mongoDB!")

	return client
}

var Client *mongo.Client = DbInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("golangApi").Collection(collectionName)

	return collection
}
