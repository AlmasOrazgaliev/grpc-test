package store

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Database struct {
	Collection *mongo.Collection
	Client     *mongo.Client
}

func NewDatabase() (database *Database, err error) {
	clientOpt := options.Client().ApplyURI("mongodb://mongodb:27017")

	client, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		log.Fatal(err)
	}

	database = &Database{
		Client: client,
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return
}
