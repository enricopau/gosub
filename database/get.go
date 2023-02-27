package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mail gets the document for the specified mail address in the specified collection.
func Mail(ctx context.Context, mail string, coll *mongo.Collection) (*MailEntry, error) {
	var result *MailEntry
	err := coll.FindOne(ctx, bson.D{{Key: "mail", Value: mail}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ID gets the document for the specified ID in the specified collection.
func ID(ctx context.Context, id string, coll *mongo.Collection) (*MailEntry, error) {
	var result *MailEntry
	err := coll.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	return result, nil
}

// Token gets the document for the specified token in the specified collection.
func Token(ctx context.Context, token string, coll *mongo.Collection) (*MailEntry, error) {
	var result *MailEntry
	err := coll.FindOne(ctx, bson.D{{Key: "token", Value: token}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	return result, nil
}

// Confirmed gets the all documents that are confirmed in the specified collection.
func Confirmed(ctx context.Context, coll *mongo.Collection) ([]MailEntry, error) {
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "confirmed", Value: false}})
	if err != nil {
		return nil, err
	}
	var entries []MailEntry
	err = cursor.All(ctx, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
