package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("step", func() {
	Description("The Step service makes it possible to view, add, modify or remove Steps of a Walkthrough.")

	HTTP(func() {
		Path("/steps")
	})

	Method("list", func() {
		Description("List all stored Steps for a given walkthrough")

		Payload(func() {
			Field(1, "id", String, "ID of Walkthrough to search for steps ")
			Required("id")
		})

		Result(StoredSteps, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("add", func() {
		Description("Add new Steps to walkthrough and return ID.")

		Payload(Steps)

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
		Description("Remove Steps from storage")

		Payload(func() {
			Field(1, "id", String, "ID of Steps to remove")
			Required("id")
		})

		Error("not_found", ElementNotFound, "Steps not found")

		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})

		GRPC(func() {
			Response(CodeOK)
		})

	})

	Method("update", func() {
		Description("Update Steps with the given IDs.")
		Payload(StoredSteps)

		HTTP(func() {
			PUT("/update")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})
