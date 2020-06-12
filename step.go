package guideme

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	db "guide.me/db"

	step "guide.me/gen/step"
)

// step service example implementation.
// The example methods log the requests and return zero values.
type stepsrvc struct {
	db     *db.Mongo
	logger *log.Logger
}

// NewStep returns the step service implementation.
func NewStep(client *mongo.Client, logger *log.Logger) (step.Service, error) {
	mongo, err := db.NewMongoClient(client)

	if err != nil {
		return nil, err
	}

	return &stepsrvc{mongo, logger}, nil

}

// List all stored Steps for a given walkthrough
func (s *stepsrvc) List(ctx context.Context, p *step.ListPayload) (res *step.StoredSteps, err error) {
	res = &step.StoredSteps{}
	s.logger.Print("step.list")
	return
}

// Add new Steps to walkthrough and return ID.
func (s *stepsrvc) Add(ctx context.Context, p *step.AddStepPayload) (res *step.ResultStep, view string, err error) {
	res = &step.ResultStep{}
	view = "default"
	s.logger.Print("step.add")

	res, err = s.db.SaveStep(p)

	if err != nil {
		log.Println(err)
		return res, view, err
	}
	return res, view, err
}

// Remove Steps from storage
func (s *stepsrvc) Remove(ctx context.Context, p *step.RemovePayload) (err error) {
	s.logger.Print("step.remove")
	return
}

// Update Steps with the given IDs.
func (s *stepsrvc) Update(ctx context.Context, p *step.StoredSteps) (err error) {
	s.logger.Print("step.update")
	return
}
