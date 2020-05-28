package database

import (
	"context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"guide.me/gen/walkthrough"
)

type WalkthroughModel struct {
	// ID is the unique id of the Walkthrough.
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Name of Walkthrough
	Name string
	// base url for your tutorial to start from
	BaseURL string
	// Status of the walkthrough [draft|published]
	Status string
	// Code to be added into an existing page to make it visible locally
	PublishedURL *string
	// ID of the organization this tutorial belongs to
	Organization string
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

func (m *Mongo) UpdateStatusWalkthrough(wtid string, wgStatus string) error {
	collection := m.getCollection(WT_COLLNAME)

	//update a Walkthrough
	id, err := primitive.ObjectIDFromHex(wtid)

	if err != nil {
		log.Println(err)
		return ErrNotFound
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{"status", wgStatus}}}}

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

func (m *Mongo) UpdateWalkthrough(walkthrough walkthrough.StoredWalkthrough) error {
	collection := m.getCollection(WT_COLLNAME)

	//update a Walkthrough
	id, err := primitive.ObjectIDFromHex(walkthrough.ID)

	if err != nil {
		log.Println(err)
		return ErrNotFound
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{"name", walkthrough.Name},
		{"baseurl", walkthrough.BaseURL},
		{"status", walkthrough.Status},
		{"publishedurl", walkthrough.PublishedURL}}}}

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

func (m *Mongo) DeleteWalkthrough(id string) error {
	collection := m.getCollection(WT_COLLNAME)

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

func (m *Mongo) LoadWalkthrough(id string) (*walkthrough.StoredWalkthrough, error) {
	collection := m.getCollection(WT_COLLNAME)

	filterId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	var element WalkthroughModel
	err = collection.FindOne(context.Background(), bson.M{"_id": filterId}).Decode(&element)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	var storedObject walkthrough.StoredWalkthrough
	storedObject = walkthrough.StoredWalkthrough{
		ID:           convertIdToString(element.ID),
		BaseURL:      element.BaseURL,
		Organization: element.Organization,
		Name:         element.Name,
		PublishedURL: element.PublishedURL,
		Status:       element.Status,
	}

	return &storedObject, err
}

func (m *Mongo) LoadAllWalkthroughs(id string) (walkthrough.StoredWalkthroughCollection, error) {
	collection := m.getCollection(WT_COLLNAME)

	// filterId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, ErrNotFound
	// }

	log.Println("id passado foi ", id)
	cursor, err := collection.Find(context.Background(), bson.M{"organization": id})

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	defer cursor.Close(context.Background())

	var listOfStoredWalkthroughs walkthrough.StoredWalkthroughCollection

	for cursor.Next(context.Background()) {
		var element WalkthroughModel

		err := cursor.Decode(&element)
		if err != nil {
			log.Println(err)
			return nil, ErrNotFound
		}

		var storedObject walkthrough.StoredWalkthrough
		storedObject = walkthrough.StoredWalkthrough{
			ID:           convertIdToString(element.ID),
			BaseURL:      element.BaseURL,
			Organization: element.Organization,
			Name:         element.Name,
			PublishedURL: element.PublishedURL,
			Status:       element.Status,
		}
		listOfStoredWalkthroughs = append(listOfStoredWalkthroughs, &storedObject)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	return listOfStoredWalkthroughs, err
}

func (m *Mongo) SaveWalkthrough(wt walkthrough.Walkthrough) (*walkthrough.StoredWalkthrough, error) {

	collection := m.getCollection(WT_COLLNAME)

	res, err := collection.InsertOne(context.Background(), wt)
	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	element := walkthrough.StoredWalkthrough{
		ID:           convertIdToString(res.InsertedID.(primitive.ObjectID)),
		BaseURL:      wt.BaseURL,
		Name:         wt.Name,
		Organization: wt.Organization,
		PublishedURL: wt.PublishedURL,
		Status:       wt.Status,
	}
	return &element, err

}

func (m *Mongo) UpdateNameWalkthrough(wtid string, wgName string) (*walkthrough.StoredWalkthrough, error) {
	collection := m.getCollection(WT_COLLNAME)
	log.Println("id passado: ", wtid, " nome passado: ", wgName)
	//update a Walkthrough
	id, err := primitive.ObjectIDFromHex(wtid)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	log.Println("deu certo converter o numero")
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{"name", wgName}}}}

	var element WalkthroughModel
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&element)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	updated := walkthrough.StoredWalkthrough{
		ID:           convertIdToString(element.ID),
		BaseURL:      element.BaseURL,
		Name:         element.Name,
		Organization: element.Organization,
		PublishedURL: element.PublishedURL,
		Status:       element.Status,
	}

	return &updated, err

}
