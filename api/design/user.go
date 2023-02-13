package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("user", func() {
	Description("")
	Security(JWTAuth)

	Method("get", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
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
