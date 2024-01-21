package database

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   sync.Once
	client *mongo.Client
)

func DBInstance() (*mongo.Client, error) {
	var err error
	once.Do(func() {
		err = godotenv.Load(".env")
		if err != nil {
			return
		}

		MongoDb := os.Getenv("MONGODB_URL")
		client, err = mongo.NewClient(options.Client().ApplyURI(MongoDb))
		if err != nil {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			return
		}

		fmt.Println("Connected to MongoDB")
	})
	return client, err
}

var Client *mongo.Client

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
