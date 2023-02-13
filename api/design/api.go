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
		})
	})
	HTTP(func() {
		Path("/api")
	})
	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("X-Requested-With")
		cors.Methods("GET", "POST")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})
})

var JWTAuth = JWTSecurity("jwt", func() {
	Description("JWT based authentication")
})
