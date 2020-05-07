package guideme

import (
	"context"
	"fmt"
	"log"

	organization "guide.me/gen/organization"
	step "guide.me/gen/step"
	"guide.me/gen/walkthrough"

	"go.mongodb.org/mongo-driver/bson"
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

func (m *Mongo) SaveWalkthrough(wt walkthrough.Walkthrough) (string, error) {

	collection := m.getCollection(WT_COLLNAME)

	res, err := collection.InsertOne(context.Background(), wt)
	if err != nil {
		log.Println(err)
		return "", ErrNotFound
	}

	return fmt.Sprintf("%v", res.InsertedID), err

}

func (m *Mongo) SaveSteps(wtSteps step.Steps) (string, error) {

	collection := m.getCollection(STEPS_COLLNAME)

	newWtId, err := primitive.ObjectIDFromHex(*wtSteps.WtID)
	if err != nil {
		log.Println(err)
		return "", ErrNotFound
	}

	stepsModel := StepsModel{
		WtID:  newWtId,
		Steps: wtSteps.Steps,
	}

	res, err := collection.InsertOne(context.Background(), stepsModel)
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

func (m *Mongo) LoadWalkthroughSteps(id string) (*step.StoredSteps, error) {
	collection := m.getCollection(STEPS_COLLNAME)

	log.Println("@@@ vou criar um novo ID")
	filterId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	log.Println("@@@ Novo ID foi sucesso")

	var element StepsModel
	err = collection.FindOne(context.Background(), bson.M{"_id": filterId}).Decode(&element)
	log.Println("@@@ Fiz a busca")
	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	log.Println("@@@ A busca foi um sucesso")

	var storedObject step.StoredSteps
	storedObject = step.StoredSteps{
		ID:    element.ID.String(),
		WtID:  element.WtID.String(),
		Steps: element.Steps,
	}

	return &storedObject, err
}

func (m *Mongo) LoadAllWalkthroughs(id string) (walkthrough.StoredWalkthroughCollection, error) {
	collection := m.getCollection(WT_COLLNAME)

	filterId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	cursor, err := collection.Find(context.Background(), bson.M{"organization": filterId})

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
			ID:           element.ID.String(),
			BaseURL:      element.BaseURL,
			Organization: element.orgID.String(),
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
		ID:           element.ID.String(),
		BaseURL:      element.BaseURL,
		Organization: element.orgID.String(),
		Name:         element.Name,
		PublishedURL: element.PublishedURL,
		Status:       element.Status,
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

func (m *Mongo) DeleteWalkthroughSteps(id string) error {
	collection := m.getCollection(STEPS_COLLNAME)

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

func (m *Mongo) UpdateWalkthrough(walkthrough walkthrough.StoredWalkthrough) error {
	collection := m.getCollection(ORG_COLLNAME)

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

func (m *Mongo) UpdateStatusWalkthrough(wtid string, wgStatus string) error {
	collection := m.getCollection(ORG_COLLNAME)

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

func (m *Mongo) UpdateWalkthroughSteps(wtSteps step.StoredSteps) error {
	collection := m.getCollection(STEPS_COLLNAME)

	//update an Organization
	id, err := primitive.ObjectIDFromHex(wtSteps.ID)

	if err != nil {
		log.Println(err)
		return ErrNotFound
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"steps", wtSteps.Steps}}}}

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
