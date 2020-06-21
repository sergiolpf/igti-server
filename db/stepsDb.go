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

type StepsModel struct {
	// walkthrough id.
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// List of steps for a given walkthrough.
	Steps []*step.StoredStep `bson:"steps,omitempty"`
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

func (m *Mongo) LoadWalkthroughSteps(id string) (*step.StoredListOfSteps, error) {
	collection := m.getCollection(STEPS_COLLNAME)

	filterId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	//log.Println("id passado por parametro eh:", id)
	var element StepsModel
	err = collection.FindOne(context.Background(), bson.M{"_id": filterId}).Decode(&element)

	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	// log.Println("resultado parece ter dados certo tamanho do array eh:", element)
	// fmt.Printf("elemento dessa chibata eh %v", element)
	// log.Println("")

	listOfSteps := make([]*step.StoredStep, 0, len(element.Steps))

	for _, myStep := range element.Steps {
		listOfSteps = append(listOfSteps, &step.StoredStep{
			ID:         myStep.ID,
			Action:     myStep.Action,
			StepNumber: myStep.StepNumber,
			Target:     myStep.Target,
			Content:    myStep.Content,
			Placement:  myStep.Placement,
			Title:      myStep.Title,
		})
	}

	response := step.StoredListOfSteps{
		WtID:  id,
		Steps: listOfSteps,
	}

	return &response, err
}

func (m *Mongo) DeleteWalkthroughStep(wtId string, id string) error {
	collection := m.getCollection(STEPS_COLLNAME)

	filterId, err := primitive.ObjectIDFromHex(wtId)

	if err != nil {
		log.Println(err)
		return ErrNotFound
	}

	filter := bson.D{{"_id", filterId}}
	update := bson.D{{"$pull", bson.D{{"steps", bson.D{{"id", id}}}}}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Println(err)
		return ErrNotFound
	}

	if result.MatchedCount != 0 {
		log.Println("matched and replaced an existing document")
	} else {
		log.Println("nothing happened!")
	}

	return nil

}

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
