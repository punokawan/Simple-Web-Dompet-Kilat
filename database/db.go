package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func ConnectDB() {
	MONGO_URI := strings.Join(
		[]string{
			"mongodb://",
			os.Getenv("MONGO_USERNAME"),
			":",
			os.Getenv("MONGO_PASSWORD"),
			"@",
			os.Getenv("MONGO_HOST"),
			":",
			os.Getenv("MONGO_PORT"),
		},
		"")

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DB")),
	}
}
