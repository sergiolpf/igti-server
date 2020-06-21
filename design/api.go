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

var ElementNotFound = Type("ElementNotFound", func() {
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("Element 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing element")
	Required("message", "id")

})
