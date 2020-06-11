package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	step "guide.me/gen/step"
)

type StepsModel struct {
	// ID is the unique id of the Steps.
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// List of steps for a given walkthrough.
	Steps []*step.Step
}

func (m *Mongo) SaveSteps(wtSteps step.Steps) (string, error) {

	collection := m.getCollection(STEPS_COLLNAME)

	newWtId, err := primitive.ObjectIDFromHex(*wtSteps.WtID)
	if err != nil {
		log.Println(err)
		return "", ErrNotFound
	}

	stepsModel := StepsModel{
		ID:    newWtId,
		Steps: wtSteps.Steps,
	}

	res, err := collection.InsertOne(context.Background(), stepsModel)
	if err != nil {
		log.Println(err)
		return "", ErrNotFound
	}

	return fmt.Sprintf("%v", res.InsertedID), err

}

func (m *Mongo) LoadWalkthroughSteps(id string) (*step.StoredSteps, error) {
	collection := m.getCollection(STEPS_COLLNAME)

	filterId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	var element StepsModel
	err = collection.FindOne(context.Background(), bson.M{"_id": filterId}).Decode(&element)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}

	var storedObject step.StoredSteps
	listOfSteps := make([]*step.Step, 0, len(element.Steps))

	for _, myStep := range element.Steps {
		listOfSteps = append(listOfSteps, &step.Step{
			Action:   myStep.Action,
			Sequence: myStep.Sequence,
			Targetid: myStep.Targetid,
			Type:     myStep.Type,
			Value:    myStep.Value,
		})
	}

	storedObject = step.StoredSteps{
		ID:    convertIdToString(element.ID),
		WtID:  convertIdToString(element.ID),
		Steps: listOfSteps,
	}

	return &storedObject, err

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
