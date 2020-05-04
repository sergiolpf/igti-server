package guideme

import (
	"context"
	"fmt"
	"log"

	organization "guide.me/gen/organization"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ErrNotFound is the error returned when attempting to load a record that does
// not exist.
var ErrNotFound = fmt.Errorf("missing record")

// COLLNAME Collection name
const ORG_COLLNAME = "organizations"

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

// Save writes the record to the DB and returns the corresponding new ID.
// data must contain a value that can be marshaled by the encoding/json package.
func (m *Mongo) SaveOneOrganization(org organization.Organization) (string, error) {

	collection := m.getCollection(ORG_COLLNAME)

	res, err := collection.InsertOne(context.Background(), org)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Adicionado ao banco com sucesso!!!")

	return fmt.Sprintf("%v", res.InsertedID), err

}

func (m *Mongo) LoadAllOrganizations() (organization.StoredOrganizationCollection, error) {
	collection := m.getCollection(ORG_COLLNAME)

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// var orgs []bson.M

	// if err = cursor.All(context.Background(), &orgs); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(orgs)

	var listOfStoredOrgs organization.StoredOrganizationCollection
	for cursor.Next(context.Background()) {
		fmt.Println("cursor!!!")
		var element organization.StoredOrganization

		err := cursor.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("decode foi ok")
		fmt.Println(element)

		listOfStoredOrgs = append(listOfStoredOrgs, &element)
	}
	fmt.Println(listOfStoredOrgs)

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return listOfStoredOrgs, err
}
