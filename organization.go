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
	res = &organization.StoredOrganization{}
	view = "default"
	s.logger.Print("organization.show")
	return
}

// Add new bottle and return its ID.
func (s *organizationsrvc) Add(ctx context.Context, p *organization.Organization) (res string, err error) {
	s.logger.Print("organization.add")

	res, err = s.db.SaveOneOrganization(*p)
	if err != nil {
		log.Fatal(err)
	}
	return res, err

}

// Remove Organization from storage
func (s *organizationsrvc) Remove(ctx context.Context, p *organization.RemovePayload) (err error) {
	s.logger.Print("organization.remove")
	return
}

// Add n number of Organizations and return their IDs. This is a multipart
// request and each part has field name 'organization' and contains the encoded
// organization info to be added.
func (s *organizationsrvc) MultiAdd(ctx context.Context, p []*organization.Organization) (res []string, err error) {
	s.logger.Print("organization.multi_add")
	return
}

// Update Organizations with the given IDs. This is a multipart request and
// each part has field name 'organizations' and contains the encoded
// Organizations info to be updated. The IDs in the query parameter is mapped
// to each part in the request.
func (s *organizationsrvc) MultiUpdate(ctx context.Context, p *organization.MultiUpdatePayload) (err error) {
	s.logger.Print("organization.multi_update")
	return
}
