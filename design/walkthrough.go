package design

import . "goa.design/goa/v3/dsl"

var _ = Service("walkthrough", func() {
	Description("The walkthrough service makes it possible to view, add, modify or remove walkthroughs.")

	HTTP(func() {
		Path("/walkthrough")
	})

	Method("list", func() {
		Description("List all stored walkthrough for a given organization")

		Payload(func() {
			Field(1, "id", String, "ID of Organization to search for ")
			Required("id")
		})

		Result(CollectionOf(StoredWalkthrough), func() {
			View("tiny")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show Walkthrough by ID")
		Payload(func() {
			Field(1, "id", String, "ID of the Walkthrough to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})

		Result(StoredWalkthrough)

		Error("not_found", ElementNotFound, "Walkthrough not found")

		HTTP(func() {
			GET("/show/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})

		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("add", func() {
		Description("Add new Tutorial and return its ID.")

		Payload(Walkthrough)

		Result(StoredWalkthrough)

		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove Walkthrough from storage")

		Payload(func() {
			Field(1, "id", String, "ID of Walkthrough to remove")
			Required("id")
		})

		Error("not_found", ElementNotFound, "Walkthrough not found")

		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})

		GRPC(func() {
			Response(CodeOK)
		})

	})

	Method("update", func() {
		Description("Update Walkthrough with the given IDs.")
		Payload(StoredWalkthrough)

		HTTP(func() {
			PUT("/update")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("rename", func() {
		Description("Rename Walkthrough with the given IDs.")
		Payload(func() {
			Field(1, "id", String, "ID of Walkthrough to be renamed")
			Field(2, "name", String, "New Name to the walkthrough")
			Required("id", "name")
		})

		Result(StoredWalkthrough, func() {
			View("tiny")
		})

		HTTP(func() {
			PUT("/rename")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("publish", func() {
		Description("Publishes Walkthrough with the given IDs.")
		Payload(func() {
			Field(1, "id", String, "ID of Walkthrough to be published")
			Required("id")
		})

		HTTP(func() {
			PUT("/publish/{id}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

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
