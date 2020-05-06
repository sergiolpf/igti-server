package guideme

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	organization "guide.me/gen/organization"
)

// organization service example implementation.
// The example methods log the requests and return zero values.
type organizationsrvc struct {
	db     *Mongo
	logger *log.Logger
}

// NewOrganization returns the organization service implementation.
func NewOrganization(client *mongo.Client, logger *log.Logger) (organization.Service, error) {
	// Setup Database
	mongo, err := NewMongoClient(client)

	if err != nil {
		return nil, err
	}

	return &organizationsrvc{mongo, logger}, nil
}

// List all stored Organizations
func (s *organizationsrvc) List(ctx context.Context) (res organization.StoredOrganizationCollection, err error) {
	s.logger.Print("organization.list")
	res, err = s.db.LoadAllOrganizations()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Show Organization by ID
func (s *organizationsrvc) Show(ctx context.Context, p *organization.ShowPayload) (res *organization.StoredOrganization, view string, err error) {
	s.logger.Print("organization.show")
	res = &organization.StoredOrganization{}
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.db.LoadOrganization(p.ID)
	if err == ErrNotFound {

		return nil, view, &ElementNotFound{
			Message: err.Error(),
			ID:      p.ID,
		}
	}
	return res, view, nil
}

// Add new bottle and return its ID.
func (s *organizationsrvc) Add(ctx context.Context, p *organization.Organization) (res string, err error) {
	s.logger.Print("organization.add")

	res, err = s.db.SaveOrganization(*p)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res, err
}

// Remove Organization from storage
func (s *organizationsrvc) Remove(ctx context.Context, p *organization.RemovePayload) (err error) {
	s.logger.Print("organization.Remove")
	err = s.db.DeleteOrganization(p.ID)

	return err
}

// Update organization with the given IDs.
func (s *organizationsrvc) Update(ctx context.Context, p *organization.StoredOrganization) (err error) {
	s.logger.Print("organization.update")
	err = s.db.UpdateOrganization(*p)
	return
}
