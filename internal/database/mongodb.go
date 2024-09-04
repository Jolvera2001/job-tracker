package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbName = "Main"
)

var MongoClient *mongo.Client

func ConnectToMongoDB() error {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not present...")
	}
	
	uri := os.Getenv("MDB_STR")
	log.Println(uri)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	MongoClient = client

	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatalln("MongoDB Client not initialized")
	}

	return MongoClient.Database(DbName).Collection(collectionName)
}
