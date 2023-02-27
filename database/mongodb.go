package database

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

// example for documents:
// {
// 	"_id" : ObjectId("XYZ"),
// 	"mail" : "example@example.org",
// 	"token" : "12345",
// 	"confirmed" : false
// }

type MailEntry struct {
	ID           primitive.ObjectID `bson:"_id",omitempty`
	Mail         string             `bson:"mail",omitempty`
	Token        string             `bson:"token",omitempty`
	Confirmed    bool               `bson:"confirmed",omitempty`
	TimeCreation time.Time          `bson:"timestamp",omitempty`
}

func (e *MailEntry) Marshal() ([]byte, error) {
	eBytes, err := bson.Marshal(e)
	if err != nil {
		return nil, err
	}
	return eBytes, nil
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

func CreateTestDataset(n int) (entries []*MailEntry) {
	r := rand.Intn(99999)
	for i := 0; i < n; i++ {
		entry := *&MailEntry{
			ID:           primitive.NewObjectID(),
			Mail:         fmt.Sprintf("%d@test.test", r),
			Token:        "0000",
			Confirmed:    false,
			TimeCreation: time.Now(),
		}
		entries = append(entries, &entry)
	}
	return entries
}
