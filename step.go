package guideme

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	step "guide.me/gen/step"
)

// step service example implementation.
// The example methods log the requests and return zero values.
type stepsrvc struct {
	db     *Mongo
	logger *log.Logger
}

// NewStep returns the step service implementation.
func NewStep(client *mongo.Client, logger *log.Logger) (step.Service, error) {

	// Setup Database
	mongo, err := NewMongoClient(client)

	if err != nil {
		return nil, err
	}

	return &stepsrvc{mongo, logger}, nil
}

// List all stored walkthrough for a given organization
func (s *stepsrvc) List(ctx context.Context, p *step.ListPayload) (res step.StoredWalkthroughCollection, err error) {
	s.logger.Print("step.list")
	return
}

// Show Walkthrough by ID
func (s *stepsrvc) Show(ctx context.Context, p *step.ShowPayload) (res *step.StoredWalkthrough, view string, err error) {
	res = &step.StoredWalkthrough{}
	view = "default"
	s.logger.Print("step.show")
	return
}

// Add new Tutorial and return its ID.
func (s *stepsrvc) Add(ctx context.Context, p *step.Walkthrough) (res string, err error) {
	s.logger.Print("step.add")
	return
}

// Remove Walkthrough from storage
func (s *stepsrvc) Remove(ctx context.Context, p *step.RemovePayload) (err error) {
	s.logger.Print("step.remove")
	return
}

// Update Walkthrough with the given IDs.
func (s *stepsrvc) Update(ctx context.Context, p *step.StoredWalkthrough) (err error) {
	s.logger.Print("step.update")
	return
}

// Publishes Walkthrough with the given IDs.
func (s *stepsrvc) Publish(ctx context.Context, p *step.PublishPayload) (err error) {
	s.logger.Print("step.publish")
	return
}
