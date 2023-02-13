package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("card", func() {
	Description("")
	// 	Error("no_criteria", String, "Missing criteria")

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
	})
})
