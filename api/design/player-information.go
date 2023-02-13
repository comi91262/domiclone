package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("player-information", func() {
	Description("")
	Security(JWTAuth)

	Method("create", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
		})

		Result(String)

		HTTP(func() {
			POST("/player-informations")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
		})

		HTTP(func() {
			DELETE("/player-informations")
			Response(StatusNoContent)
		})
	})

	Method("get_coins", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
		})

		Result(Int)
		HTTP(func() {
			GET("/coins")
			Response(StatusOK)
		})
	})

	Method("get_victory_points", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
		})

		Result(Int)
		HTTP(func() {
			GET("/victory-points")
			Response(StatusOK)
		})
	})

	Method("get_decks", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
		})

		Result(Int)
		HTTP(func() {
			GET("/decks")
			Response(StatusOK)
		})
	})

	Method("get_discards", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
		})

		Result(Int)
		HTTP(func() {
			GET("/discards")
			Response(StatusOK)
		})
	})

	Method("get_hands", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
		})

		Result(Int)
		HTTP(func() {
			GET("/hands")
			Response(StatusOK)
		})
	})

	Method("get_play_area", func() {
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Required("token")
		})

		Result(Int)
		HTTP(func() {
			GET("/play-areas")
			Response(StatusOK)
		})
	})
})
