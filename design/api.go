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
		Services("organization", "walkthrough", "step", "swagger")

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

var StoredSteps = ResultType("application/vnd.goa.guide.me.stored-steps", func() {
	Description("A StoredStep describes all the Steps retrieved by the Steps service.")
	Reference(Steps)
	TypeName("StoredSteps")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the Step.", func() {
			Example("123abc")
			Meta("rpc:tag", "1")

		})
		Field(2, "wtId")
		Field(3, "steps")
	})
	View("default", func() {
		Attribute("id")
		Attribute("wtId")
		Attribute("steps")

	})
	View("tiny", func() {
		Attribute("id")
		Attribute("wtId")
		Attribute("steps")
	})
	Required("id", "wtId", "steps")
})

var Steps = Type("Steps", func() {
	Description("Steps describes the entire chain of steps to be displayed as part of the walkthrough.")
	Attribute("wtId", String, "The id of the Walkthrough those steps belong to.", func() {
		Example("abc234235")
		Meta("rpc:tag", "1")
	})

	Attribute("steps", ArrayOf(Step), "List of steps for a given walkthrough.", func() {
		Meta("rpc:tag", "2")
	})
})

var Step = Type("Step", func() {
	Description("Step describes the basic details of your tutorials.")
	Attribute("targetid", String, "A string representing the HTML ID of an element", func() {
		Example("")
		Meta("rpc:tag", "1")
	})
	Attribute("type", String, "The type of step to be used", func() {
		Enum("text", "picture")
		Default("text")
		Example("text")
		Meta("rpc:tag", "2")
	})
	Attribute("value", String, "The content of the message to be displayed", func() {
		Example("This dropdown contains values from the list of status, for our scenario we want to chose 'active'")
		Meta("rpc:tag", "3")
	})
	Attribute("sequence", Int32, "The number in the sequence that the step belongs to.", func() {
		Meta("rpc:tag", "4")
	})
	Attribute("action", String, "What action should trigger the next step", func() {
		Enum("click", "next", "end")
		Meta("rpc:tag", "5")
	})

	Required("targetid", "type", "value", "sequence", "action")

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

var ElementNotFound = Type("ElementNotFound", func() {
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("Element 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing element")
	Required("message", "id")

})
