package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Use("/connect", func(c *fiber.Ctx) error {
		// Hier wordt gekeken of de request de "Upgrade naar WebSocket" headers heeft.
		if websocket.IsWebSocketUpgrade(c) {
			// |-> Zo ja, ga dan naar de volgende /connect endpoint
			return c.Next()
		}
		// |-> Zo niet, stuur dan een "426 Upgrade Required"
		return fiber.ErrUpgradeRequired
	})

	app.Get("/connect", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			// Genereer een willekeurig getal.
			// |-> Dit is alleen ter demonstratie
			mt = 1
			msg = []byte("string")

			// Stuur berichten [commando's] via de WebSocket naar de client
			if err = c.WriteMessage(mt, msg); err != nil {
				break
			}
		}

		// for i := 0; i < 20; i++ {
		// 	mt = 1
		// 	msg = []byte("string")

		// 	// Stuur berichten [commando's] via de WebSocket naar de client
		// 	if err = c.WriteMessage(mt, msg); err != nil {
		// 		break
		// 	}
		// }

	}))

	log.Fatal(app.Listen(":3000"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
