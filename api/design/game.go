package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("game", func() {
	Description("")
	// 	Error("no_criteria", String, "Missing criteria")
	Method("get", func() {
		Payload(func() {
			Attribute("id", Int, "Game ID", func() {})
			Required("id")
		})
		Result(String)
		HTTP(func() {
			GET("/games/{id}")
			Response(StatusOK)
		})
	})

	Method("create", func() {
		Result(Int)
		HTTP(func() {
			POST("/games")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Payload(func() {
			Attribute("id", Int, "Game ID", func() {})
			Required("id")
		})
		HTTP(func() {
			DELETE("/games/{id}")
			Response(StatusNoContent)
		})
	})

	Method("get_supplies", func() {
		Payload(func() {
			Attribute("id", Int, "Game ID", func() {})
			Required("id")
		})
		Result(String)
		HTTP(func() {
			GET("/games/{id}/supplies")
			Response(StatusOK)
		})
	})

	Method("get_trashes", func() {
		Payload(func() {
			Attribute("id", Int, "Game ID", func() {})
			Required("id")
		})
		Result(String)
		HTTP(func() {
			GET("/games/{id}/trashes")
			Response(StatusOK)
		})
	})
})
