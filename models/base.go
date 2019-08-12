package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var dbClient *mongo.Client

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Vc346:Victobias1@gmds-fa1px.mongodb.net/test?w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	dbClient = client
}

var GetClient = func() *mongo.Client {
	return dbClient
}
