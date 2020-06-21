package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ErrNotFound is the error returned when attempting to load a record that does
// not exist.
var ErrNotFound = fmt.Errorf("missing record")

// COLLNAME Collection name
const ORG_COLLNAME = "organizations"

const WT_COLLNAME = "walkthroughs"

const STEPS_COLLNAME = "steps"

// DBNAME Database name
const DBNAME = "arahi"

// Mongo is the database driver.
type Mongo struct {
	// client is the Bolt client.
	db *mongo.Database
}

// NewBoltDB creates a Bolt DB database driver given an underlying client.
func NewMongoClient(client *mongo.Client) (*Mongo, error) {
	err := client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		err = client.Connect(context.Background())

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		defer client.Disconnect(context.Background())
	}

	return &Mongo{client.Database(DBNAME)}, nil

}

func (m *Mongo) getCollection(collectionName string) *mongo.Collection {

	collection := m.db.Collection(collectionName)

	return collection
}

func convertIdToString(id primitive.ObjectID) string {

	newId, err := id.MarshalJSON()
	if err != nil {
		log.Println(err)
		return ""
	}
	stringId := strings.Replace(string(newId), "\"", "", -1)
	return stringId
}
