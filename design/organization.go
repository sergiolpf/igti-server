package design

import . "goa.design/goa/v3/dsl"

var _ = Service("organization", func() {
	Description("The Organization service makes it possible to view, add or remove Organizations.")

	HTTP(func() {
		Path("/organization")
	})

	Method("list", func() {
		Description("List all stored Organizations")

		Result(CollectionOf(StoredOrganization), func() {
			View("tiny")
		})

		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show Organization by ID")
		Payload(func() {
			Field(1, "id", String, "ID of the Organization to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})

		Result(StoredOrganization)

		Error("not_found", ElementNotFound, "Organization not found")

		HTTP(func() {
			GET("/{id}")
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
		Description("Add new bottle and return its ID.")

		Payload(Organization)

		Result(String)

		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove Organization from storage")

		Payload(func() {
			Field(1, "id", String, "ID of Organization to remove")
			Required("id")
		})

		Error("not_found", ElementNotFound, "Organization not found")

		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})

		GRPC(func() {
			Response(CodeOK)
		})

	})

	Method("update", func() {
		Description("Update organization with the given IDs.")
		Payload(StoredOrganization)

		HTTP(func() {
			PUT("/update")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
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
