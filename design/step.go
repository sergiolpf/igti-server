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

		Result(StoredListOfSteps, func() {
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

		Payload(AddStepPayload)

		Result(ResultStep)

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
			Attribute("wtId", String, "Id of the Walkthrough", func() {
				Meta("rpc:tag", "1")
			})
			Field(2, "id", String, "ID of the step to be remove")
			Required("wtId", "id")
		})

		Error("not_found", ElementNotFound, "Steps not found")

		HTTP(func() {
			DELETE("/")
			Response(StatusNoContent)
		})

		GRPC(func() {
			Response(CodeOK)
		})

	})

	// Method("update", func() {
	// 	Description("Update Steps with the given IDs.")
	// 	Payload(StoredSteps)

	// 	HTTP(func() {
	// 		PUT("/update")
	// 		Response(StatusNoContent)
	// 	})
	// 	GRPC(func() {
	// 		Response(CodeOK)
	// 	})
	// })
})

var Step = Type("Step", func() {
	Description("Step describes the basic details of your tutorials.")
	Attribute("title", String, "Title for the given step", func() {
		Example("Click here to make it work!")
		Meta("rpc:tag", "1")
	})
	Attribute("target", String, "Unique html if for the target", func() {
		Meta("rpc:tag", "2")
	})
	Attribute("stepNumber", Int32, "The number in the sequence that the step belongs to", func() {
		Meta("rpc:tag", "3")
	})
	Attribute("placement", String, "Where the popup will be anchored, left, right, top or buttom.", func() {
		Enum("left", "right", "top", "buttom")
		Default("right")
		Meta("rpc:tag", "4")
	})
	Attribute("content", String, "The content of the message to be displayed", func() {
		Example("This dropdown contains values from the list of status, for our scenario we want to chose 'active'")
		Meta("rpc:tag", "5")
	})
	Attribute("action", String, "What action should trigger the next step", func() {
		Enum("click", "next", "end")
		Default("next")
		Meta("rpc:tag", "6")
	})

	Required("title", "target", "stepNumber", "placement", "content", "action")

})

var StoredStep = ResultType("application/vnd.goa.guide.me.stored-step", func() {
	Description("A StoredStep describes a step returned from the database.")
	TypeName("StoredStep")
	Reference(Step)

	Attributes(func() {
		Attribute("id", String, "Unique id to this step", func() {

			Meta("rpc:tag", "1")
		})
		Field(2, "title")
		Field(3, "target")
		Field(4, "stepNumber")
		Field(5, "placement")
		Field(6, "content")
		Field(7, "action")

	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("target")
		Attribute("stepNumber")
		Attribute("placement")
		Attribute("content")
		Attribute("action")

	})
	View("tiny", func() {
		Attribute("id")
		Attribute("title")
		Attribute("target")
		Attribute("stepNumber")
		Attribute("content")
	})
	Required("id", "title", "target", "stepNumber", "content")
})

var StoredListOfSteps = ResultType("application/vnd.goa.guide.me.stored-listof-steps", func() {
	Description("A StoredListOfSteps describes all the Steps retrieved by the Steps service.")
	Reference(StoredStep)
	TypeName("StoredListOfSteps")

	Attributes(func() {
		Attribute("wtId", String, "ID is the unique id of the Walkthrough.", func() {
			Example("123abc")
			Meta("rpc:tag", "1")

		})
		Attribute("steps", ArrayOf(StoredStep), "List of Stored steps", func() {

			Meta("rpc:tag", "2")
		})
	})

	View("default", func() {
		Attribute("wtId")
		Attribute("steps")
	})
	Required("wtId", "steps")
})

var AddStepPayload = Type("AddStepPayload", func() {
	Description("Payload to be used the POST method to add a new step for a walkthrough.")
	Attribute("wtId", String, "Id of the walkthrough to have a step added to", func() {

		Meta("rpc:tag", "1")
	})
	Attribute("step", Step, "step to be added", func() {

		Meta("rpc:tag", "2")
	})
})

var ResultStep = ResultType("application/vnd.goa.guide.me.result-steps", func() {
	Description("A ResultStep describes the added/modified step for a given walkthrough.")
	TypeName("ResultStep")

	Attributes(func() {
		Attribute("wtId", String, "Id of the walkthrough to have a step added to", func() {

			Meta("rpc:tag", "1")
		})
		Attribute("step", StoredStep, "Modified step", func() {

			Meta("rpc:tag", "2")
		})

	})
	View("default", func() {
		Attribute("wtId")
		Attribute("step")

	})
	View("tiny", func() {
		Attribute("wtId")
		Attribute("step")
	})
	Required("wtId", "step")
})
