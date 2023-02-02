package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mail gets the document for the specified mail address in the specified collection.
func Mail(mail string, coll *mongo.Collection) bson.M {
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "mail", Value: mail}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	return result
}

// ID gets the document for the specified ID in the specified collection.
func ID(id string, coll *mongo.Collection) bson.M {
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	return result
}

// Token gets the document for the specified token in the specified collection.
func Token(token string, coll *mongo.Collection) bson.M {
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "token", Value: token}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	return result
}

// Confirmed gets the all documents that are confirmed in the specified collection.
func Confirmed(coll *mongo.Collection) ([]MailEntry, error) {
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "confirmed", Value: false}})
	if err != nil {
		return nil, err
	}
	var entries []MailEntry
	err = cursor.All(context.TODO(), &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
