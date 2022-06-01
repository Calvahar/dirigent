package main

import (
	"fmt"
	"log"

	iki "github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

func main() {
	// Alle verbonden clients worden hier opgeslagen.
	clients := make(map[string]string)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/test", func(c *fiber.Ctx) error {
		fmt.Println(clients)
		return c.Send([]byte("test"))
	})

	// Middleware voor de eerste GET request naar de WebSocket endpoint
	app.Use("/connect", func(c *fiber.Ctx) error {
		/* Controleer of er WebSocket headers aanwezig zijn in de GET request
		Connection: Upgrade
		Upgrade: websocket
		*/

		if len(c.Params("key")) < 1 {
			return fiber.ErrUnauthorized
		}
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			c.Locals("UniqueId")
			// Ga door naar de volgende /connect
			return c.Next()
		}
		// Return een 426 status
		return fiber.ErrUpgradeRequired
	})

	app.Get("/connect", iki.New(func(conn *iki.Websocket) {
		userId := conn.Params("key")

		clients[userId] = conn.UUID

		conn.SetAttribute("user_id", userId)

		conn.Broadcast([]byte("Broadcast"), true)

		conn.Emit([]byte("Emit"))
	}))

	// Events
	iki.On(iki.EventDisconnect, func(ep *iki.EventPayload) {
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
	})

	iki.On(iki.EventClose, func(ep *iki.EventPayload) {
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
	})

	log.Fatal(app.Listen(":3000"))
}
