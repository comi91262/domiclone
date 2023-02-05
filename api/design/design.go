package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("domilike", func() {
	Title("Domilike Service")
	Description("")
	Server("domilike", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("card", func() {
	Description("")

	Method("get", func() {
		Payload(func() {
			Attribute("id", Int, "Card ID", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})

		Result(String)

		HTTP(func() {
			GET("/cards/{id}")
			Response(StatusOK)
		})

		GRPC(func() {
		})
	})
})
