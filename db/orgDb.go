package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	organization "guide.me/gen/organization"
)

type OrganizationModel struct {
	// ID is the unique id of the Organization.
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Name of Organization
	Name string
	// Company website URL
	URL string
}

// Save writes the record to the DB and returns the corresponding new ID.
// data must contain a value that can be marshaled by the encoding/json package.
func (m *Mongo) SaveOrganization(org organization.Organization) (string, error) {

	collection := m.getCollection(ORG_COLLNAME)

	res, err := collection.InsertOne(context.Background(), org)
	if err != nil {
		log.Println(err)
		return "", ErrNotFound
	}

	return fmt.Sprintf("%v", res.InsertedID), err

}

func (m *Mongo) LoadAllOrganizations() (organization.StoredOrganizationCollection, error) {
	collection := m.getCollection(ORG_COLLNAME)

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	defer cursor.Close(context.Background())

	var listOfStoredOrgs organization.StoredOrganizationCollection

	for cursor.Next(context.Background()) {
		var element OrganizationModel

		err := cursor.Decode(&element)
		if err != nil {
			log.Println(err)
			return nil, ErrNotFound
		}

		var storedObject organization.StoredOrganization
		storedObject = organization.StoredOrganization{
			ID:   element.ID.String(),
			Name: element.Name,
			URL:  element.URL,
		}
		listOfStoredOrgs = append(listOfStoredOrgs, &storedObject)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	return listOfStoredOrgs, err
}

func (m *Mongo) LoadOrganization(id string) (*organization.StoredOrganization, error) {
	collection := m.getCollection(ORG_COLLNAME)

	filterId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	var element OrganizationModel
	err = collection.FindOne(context.Background(), bson.M{"_id": filterId}).Decode(&element)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	var storedObject organization.StoredOrganization
	storedObject = organization.StoredOrganization{
		ID:   element.ID.String(),
		Name: element.Name,
		URL:  element.URL,
	}

	return &storedObject, err
}

func (m *Mongo) DeleteOrganization(id string) error {
	collection := m.getCollection(ORG_COLLNAME)

	filterId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return ErrNotFound
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": filterId})

	if err != nil {
		log.Println(err)
		return ErrNotFound
	}
	return nil

}

func (m *Mongo) UpdateOrganization(org organization.StoredOrganization) error {
	collection := m.getCollection(ORG_COLLNAME)

	//update an Organization
	id, err := primitive.ObjectIDFromHex(org.ID)

	if err != nil {
		log.Println(err)
		return ErrNotFound
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", org.Name}, {"url", org.URL}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return ErrNotFound
	}

	if result.MatchedCount != 0 {
		log.Println("matched and replaced an existing document")
	}

	return err

}
