package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(address string, port int64) *mongo.Database {
	//clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", address, port))
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", address, port))

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("microblog")

	return database
}
