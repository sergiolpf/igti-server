package guideme

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	walkthrough "guide.me/gen/walkthrough"
)

type WalkthroughStatusType string

const (
	Draft     WalkthroughStatusType = "draft"
	Published                       = "published"
	Removed                         = "removed"
)

// walkthrough service example implementation.
// The example methods log the requests and return zero values.
type walkthroughsrvc struct {
	db     *Mongo
	logger *log.Logger
}

// NewWalkthrough returns the walkthrough service implementation.
func NewWalkthrough(client *mongo.Client, logger *log.Logger) (walkthrough.Service, error) {
	// Setup Database
	mongo, err := NewMongoClient(client)

	if err != nil {
		return nil, err
	}

	return &walkthroughsrvc{mongo, logger}, nil
}

// List all stored walkthrough for a given organization
func (s *walkthroughsrvc) List(ctx context.Context, p *walkthrough.ListPayload) (res walkthrough.StoredWalkthroughCollection, err error) {
	s.logger.Print("walkthrough.list")
	res, err = s.db.LoadAllWalkthroughs(p.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Show Walkthrough by ID
func (s *walkthroughsrvc) Show(ctx context.Context, p *walkthrough.ShowPayload) (res *walkthrough.StoredWalkthrough, view string, err error) {
	res = &walkthrough.StoredWalkthrough{}
	view = "default"
	s.logger.Print("walkthrough.show")
	res = &walkthrough.StoredWalkthrough{}
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.db.LoadWalkthrough(p.ID)
	if err == ErrNotFound {

		return nil, view, &walkthrough.ElementNotFound{
			Message: err.Error(),
			ID:      p.ID,
		}
	}
	return res, view, nil

}

// Add new Tutorial and return its ID.
func (s *walkthroughsrvc) Add(ctx context.Context, p *walkthrough.Walkthrough) (res string, err error) {
	s.logger.Print("walkthrough.add")

	res, err = s.db.SaveWalkthrough(*p)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res, err
}

// Remove Walkthrough from storage
func (s *walkthroughsrvc) Remove(ctx context.Context, p *walkthrough.RemovePayload) (err error) {
	s.logger.Print("walkthrough.remove")
	err = s.db.DeleteWalkthrough(p.ID)

	return err
}

// Update Walkthrough with the given IDs.
func (s *walkthroughsrvc) Update(ctx context.Context, p *walkthrough.StoredWalkthrough) (err error) {
	s.logger.Print("walkthrough.update")
	err = s.db.UpdateWalkthrough(*p)
	return err
}

// Publishes Walkthrough with the given IDs.
func (s *walkthroughsrvc) Publish(ctx context.Context, p *walkthrough.PublishPayload) (err error) {
	s.logger.Print("walkthrough.publish")
	err = s.db.UpdateStatusWalkthrough(p.ID, "published")

	return err
}
