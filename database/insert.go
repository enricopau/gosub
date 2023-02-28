package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddEntries(ctx context.Context, entries []*MailEntry, coll *mongo.Collection) error {
	if entries == nil {
		return fmt.Errorf("mail entry cannot be nil")
	}
	var docs []interface{}
	for _, e := range entries {
		docs = append(docs, e)
	}
	_, err := coll.InsertMany(ctx, docs, &options.InsertManyOptions{})
	if err != nil {
		return err
	}
	return nil
}

func AddEntry(ctx context.Context, entry *MailEntry, coll *mongo.Collection) (id string, err error) {
	if entry == nil {
		return "", fmt.Errorf("mail entry cannot be nil")
	}
	insRes, err := coll.InsertOne(ctx, entry, &options.InsertOneOptions{})
	if err != nil {
		return "", err
	}
	return insRes.InsertedID.(primitive.ObjectID).String(), nil
}
