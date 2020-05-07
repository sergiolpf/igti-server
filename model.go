package guideme

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrganizationModel struct {
	// ID is the unique id of the Organization.
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Name of Organization
	Name string
	// Company website URL
	URL string
}

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
	orgID primitive.ObjectID `bson:"orgId,omitempty"`
}
