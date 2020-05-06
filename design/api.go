package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("guideme", func() {
	Title("Guideme Services")
	Description("HTTP service for managing the organizations and tutorials.")

	Server("guideme", func() {
		Description("guideme hosts the organizations, tutoriais and swagger services.")

		//Services("organization", "tutorial", "swagger")
		Services("organization", "walkthrough", "swagger")

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

var StoredWalkthrough = ResultType("application/vnd.goa.guide.me.stored-walkthrough", func() {
	Description("A StoredWalkthrough describes a Walkthrough retrieved by the Walkthrough service.")
	Reference(Walkthrough)
	TypeName("StoredWalkthrough")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the Walkthrough.", func() {
			Example("123abc")
			Meta("rpc:tag", "1")

		})
		Field(2, "name")
		Field(3, "baseURL")
		Field(4, "status")
		Field(5, "publishedURL")
		Field(6, "organization")

	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("baseURL")
		Attribute("status")
		Attribute("publishedURL")
		Attribute("organization")
	})
	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("baseURL")
		Attribute("organization")
	})
	Required("id", "name", "baseURL", "organization")
})

var Walkthrough = Type("Walkthrough", func() {
	Description("Walkthrough describes the basic details of your tutorials.")
	Attribute("name", String, "Name of the Tutorial", func() {
		MaxLength(100)
		Example("How to create a new process using the exception condition.")
		Meta("rpc:tag", "1")
	})
	Attribute("baseURL", String, "base url for your tutorial to start from", func() {
		Pattern(`(?i)^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
		Example("http://www.google.com/")
		Meta("rpc:tag", "2")
	})
	Attribute("status", String, "Status of the walkthrough [draft|published] ", func() {
		Default("draft")
		Enum("draft", "completed", "removed")
		Example("draft | published")
		Meta("rpc:tag", "3")
	})
	Attribute("publishedURL", String, "Code to be added into an existing page to make it visible locally", func() {
		Meta("rpc:tag", "4")
	})
	Attribute("organization", String, "ID of the organization this tutorial belongs to", func() {
		Meta("rpc:tag", "5")
	})

	Required("name", "baseURL", "organization")

})

var Organization = Type("Organization", func() {
	Description("Organization describes an Organization to be stored.")
	Attribute("name", String, "Name of Organization", func() {
		MaxLength(200)
		Example("Creating a new request in netflix!")
		Meta("rpc:tag", "1")
	})
	Attribute("url", String, "Company website URL", func() {
		Pattern(`(?i)^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
		Example("http://www.google.com/")
		Meta("rpc:tag", "2")
	})

	Required("name", "url")

})

var ElementNotFound = Type("NotFound", func() {
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("Element 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing element")
	Required("message", "id")

})
