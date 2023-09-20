package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Collection *mongo.Collection

func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27037")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database("mydb").Collection("entities")
	//defer client.Disconnect(context.Background())
}
