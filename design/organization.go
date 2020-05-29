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
