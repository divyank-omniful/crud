package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Init() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27037")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
