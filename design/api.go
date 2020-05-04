package design

import . "goa.design/goa/v3/dsl"

var _ = API("guideme", func() {
	Title("Guideme Services")
	Description("HTTP service for managing the organizations and tutorials.")

	Server("guideme", func() {
		Description("guideme hosts the organizations, tutoriais and swagger services.")

		//Services("organization", "tutorial", "swagger")
		Services("organization", "swagger")

		Host("development", func() {
			Description("Development host")
			URI("http://localhost:8000/")
			URI("grpc://localhost:8080/")
		})
	})
})

var StoredOrganization = ResultType("application/vnd.goa.guide.me.stored-organization", func() {
	Description("A StoredOrganization describes an Organization retrieved by the Organization service.")
	Reference(Organization)
	TypeName("StoredOrganization")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the Organization.", func() {
			Example("123abc")
			Meta("rpc:tag", "1")
			Meta("struct:field:type", "primitive.ObjectId", "go.mongodb.org/mongo-driver/bson/primitive")
		})
		Field(2, "name")
		Field(3, "url")

	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("url")
	})
	View("tiny", func() {
		Attribute("id")
		Attribute("name")

	})
	Required("id", "name", "url")
})

var Organization = Type("Organization", func() {
	Description("Organization describes an Organization to be stored.")
	Attribute("name", String, "Name of Organization", func() {
		MaxLength(100)
		Example("Blue's Cuvee")
		Meta("rpc:tag", "1")
	})
	Attribute("url", String, "Company website URL", func() {
		Pattern(`(?i)^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
		Example("http://www.google.com/")
		Meta("rpc:tag", "2")
	})

	Required("name", "url")

})

var OrgNotFound = Type("OrgNotFound", func() {
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("Organization 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing Organization")
	Required("message", "id")

})
