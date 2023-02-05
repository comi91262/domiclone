package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("domilike", func() {
	Title("Domilike Service")
	Description("")
	Server("domilike", func() {
		Host("localhost", func() {
			URI("http://localhost:18080")
			URI("grpc://localhost:28080")
		})
	})
	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("X-Requested-With")
		cors.Methods("GET", "POST")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
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
