package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// example for documents:
// {
// 	"_id" : ObjectId("XYZ"),
// 	"mail" : "example@example.org",
// 	"token" : "12345",
// 	"confirmed" : false
// }

type MailEntry struct {
	ID        primitive.ObjectID
	Mail      string
	Token     string
	Confirmed bool
}

const (
	dbName         = "gosub"
	collectionName = "maillist"
)

// Connect reads the URI string from the given .env file or other environment variables and connects to the mongodb client.
func Connect() (*mongo.Client, error, func(ctx context.Context) error) {
	err := godotenv.Load()
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, fmt.Errorf("no connection uri set"), nil
	}
	if err != nil {
		return nil, err, nil
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return client, err, client.Disconnect
}

// MailList is a wrapper for querying the maillist collection
// from the gosub database for the established client.
func MailListCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func AddEntry(entry *MailEntry, coll *mongo.Collection) error {
	//TODO: add insert
	return nil
}
