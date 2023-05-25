package database

import (
	"context"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const projectDirName = "golang-clean-arch-api"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func MongoDbInstance() *mongo.Client {
	loadEnv()

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

	return client
}

var MongoClient *mongo.Client = MongoDbInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("golangApi").Collection(collectionName)

	return collection
}
