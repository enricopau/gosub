package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteEntryByID removed one entry by its respective id.
func DeleteEntryByID(ctx context.Context, id primitive.ObjectID, coll *mongo.Collection) error {
	filter := bson.D{{"_id", id}}
	res, err := coll.DeleteOne(ctx, filter)
	switch {
	case err != nil:
		return err
	case res.DeletedCount == 0:
		return fmt.Errorf("no matching id found")
	default:
		return nil
	}
}

// DeleteEntriesByMail removes all occurrences of this mail address in the database.
func DeleteEntriesByMail(ctx context.Context, mail string, coll *mongo.Collection) error {
	filter := bson.D{{"mail", mail}}
	res, err := coll.DeleteMany(ctx, filter)
	switch {
	case err != nil:
		return err
	case res.DeletedCount == 0:
		return fmt.Errorf("no matching mail found")
	default:
		return nil
	}
}

// DeleteEntriesAfter finds all entries that are older than the given time.
func DeleteUnconfirmedEntriesAfter(ctx context.Context, t time.Time, coll *mongo.Collection) error {
	pt := primitive.NewDateTimeFromTime(t)
	filter := bson.D{{"$and", bson.A{
		bson.D{{"timestamp", bson.D{{"$lt", pt}}}},
		bson.D{{"confirmed", false}},
	},
	},
	}
	_, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}
