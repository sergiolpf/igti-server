package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	step "guide.me/gen/step"
)

type StepModel struct {
	// ID is the unique id of the Steps.
	ID string `bson:"id,omitempty"`
	// List of steps for a given walkthrough.
	Title string `bson:"title,omitempty"`
	// A string representing the HTML ID of an element
	Target string `bson:"target,omitempty"`
	// The number in the sequence that the step belongs to.
	StepNumber int32 `bson:"stepNumber,omitempty"`
	// The placement of the popup
	Placement string `bson:"placement,omitempty"`
	// The content of the message to be displayed
	Content string `bson:"content,omitempty"`
	// What action should trigger the next step
	Action string `bson:"action,omitempty"`
}

func (m *Mongo) SaveStep(wtSteps *step.AddStepPayload) (*step.ResultStep, error) {

	collection := m.getCollection(STEPS_COLLNAME)

	newWtId, err := primitive.ObjectIDFromHex(*wtSteps.WtID)

	defaultResultStep := step.ResultStep{}

	if err != nil {
		log.Println(err)
		return &defaultResultStep, ErrNotFound
	}

	newStepId := primitive.NewObjectID()

	stepToBeAdded := step.StoredStep{
		ID:         convertIdToString(newStepId),
		Action:     wtSteps.Step.Action,
		Content:    wtSteps.Step.Content,
		Placement:  wtSteps.Step.Placement,
		StepNumber: wtSteps.Step.StepNumber,
		Target:     wtSteps.Step.Target,
		Title:      wtSteps.Step.Title,
	}

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"_id", newWtId}}
	update := bson.D{{"$push", bson.D{{"steps", stepToBeAdded}}}}

	var updatedDocument bson.M

	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println("Added the document as it did not existed before!")
			err = nil
		}

	}

	result := step.ResultStep{
		WtID: *wtSteps.WtID,
		Step: &stepToBeAdded,
	}
	//fmt.Printf("updated document %v", updatedDocument)
	return &result, err

}

// func (m *Mongo) LoadWalkthroughSteps(id string) (*step.StoredSteps, error) {
// 	collection := m.getCollection(STEPS_COLLNAME)

// 	filterId, err := primitive.ObjectIDFromHex(id)

// 	if err != nil {
// 		log.Println(err)
// 		return nil, ErrNotFound
// 	}

// 	var element StepsModel
// 	err = collection.FindOne(context.Background(), bson.M{"_id": filterId}).Decode(&element)

// 	if err != nil {
// 		log.Println(err)
// 		return nil, ErrNotFound
// 	}

// 	var storedObject step.StoredSteps
// 	listOfSteps := make([]*step.Step, 0, len(element.Steps))

// 	for _, myStep := range element.Steps {
// 		listOfSteps = append(listOfSteps, &step.Step{

// 			Action:   myStep.Action,
// 			Sequence: myStep.Sequence,
// 			Targetid: myStep.Targetid,
// 			Type:     myStep.Type,
// 			Value:    myStep.Value,
// 		})
// 	}

// 	storedObject = step.StoredSteps{
// 		WtID:  convertIdToString(element.ID),
// 		Steps: listOfSteps,
// 	}

// 	return &storedObject, err

// }

// func (m *Mongo) DeleteWalkthroughSteps(id string) error {
// 	collection := m.getCollection(STEPS_COLLNAME)

// 	filterId, err := primitive.ObjectIDFromHex(id)

// 	if err != nil {
// 		log.Println(err)
// 		return ErrNotFound
// 	}

// 	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": filterId})

// 	if err != nil {
// 		log.Println(err)
// 		return ErrNotFound
// 	}
// 	return nil

// }

// func (m *Mongo) UpdateWalkthroughSteps(wtSteps step.StoredSteps) error {
// 	collection := m.getCollection(STEPS_COLLNAME)

// 	//update an Organization
// 	id, err := primitive.ObjectIDFromHex(wtSteps.ID)

// 	if err != nil {
// 		log.Println(err)
// 		return ErrNotFound
// 	}

// 	filter := bson.D{{"_id", id}}
// 	update := bson.D{{"$set", bson.D{{"steps", wtSteps.Steps}}}}

// 	result, err := collection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		log.Println(err)
// 		return ErrNotFound
// 	}

// 	if result.MatchedCount != 0 {
// 		log.Println("matched and replaced an existing document")
// 	}

// 	return err

// }
