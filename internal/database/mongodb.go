package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbName = "Main"
)

var MongoClient *mongo.Client

func ConnectToMongoDB() error {
	uri := os.Getenv("MDB_STR")
	if uri == "" {
		log.Fatalln("MDB_STR NOT SET!")
	}

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetTLSConfig(nil)

	// attempting connection
	log.Println("Attempting to connect to mongodb...")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	log.Println("Pinging Mongodb...")
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	log.Println("MongoDB Connection successful!")
	MongoClient = client
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatalln("MongoDB Client not initialized")
	}

	return MongoClient.Database(DbName).Collection(collectionName)
}
