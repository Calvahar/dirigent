package main

import (
	"fmt"
	"log"
	"time"

	ws "github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

type LogConfig struct {
	Connect    bool
	Disconnect bool
	Error      bool
	Logger     bool
}

func main() {
	config := LogConfig{Connect: false, Disconnect: false, Error: true, Logger: false}
	// Alle verbonden clients worden hier opgeslagen - worden verwijderd bij disconnect
	clients := make(map[string]string)

	// Nieuwe Fiber applicatie
	app := fiber.New(fiber.Config{
		AppName: "Dirigent Server v1.1.0",
	})

	/** ----------------------------------

		Middleware

	**/
	app.Use(recover.New())
	if config.Logger {
		app.Use(logger.New(logger.Config{
			// TimeZone:   "GMT+2",
			TimeFormat: "Jan-02, 15:04:05.000",
			Format:     "[${magenta}${time}${reset}] ♦ ${blue}${ip}${reset}:${cyan}${port}${reset}  ${yellow}-> ${method}to ${cyan}${url} ${reset}with${status} \n",
		}))
	}

	/** ----------------------------------

		HTTP non-keep-alive Requests

	**/
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Send([]byte("test"))
	})

	/** ----------------------------------

		WebSocket Requests

	**/
	app.Use(func(c *fiber.Ctx) error {
		// Als de request WebSocket Headers heeft, mag het door naar de volgende endpoint
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		// Anders return een 426, Upgrade required
		return fiber.ErrUpgradeRequired
	})

	/** ----------------------------------

		Events

	**/
	// Client maakt verbinding met WebSocket
	ws.On(ws.EventConnect, func(ep *ws.EventPayload) {
		if config.Connect {
			fmt.Printf("[Connect] - User: %s \n", ep.Kws.GetStringAttribute("user_id"))
		}
	})

	// Client verliest verbinding met WebSocket
	ws.On(ws.EventDisconnect, func(ep *ws.EventPayload) {
		// Client wordt uit de lijst van verbonden clients gehaald
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
		if config.Disconnect {
			fmt.Printf("[Disconnect] - User: %s \n", ep.Kws.GetStringAttribute("user_id"))
		}
	})

	// Client's verbinding krijgt een error
	ws.On(ws.EventError, func(ep *ws.EventPayload) {
		if config.Error {
			fmt.Printf("[Error] - User: %s \n", ep.Kws.GetStringAttribute("user_id"))
		}
	})

	app.Get("/ws/:id", ws.New(func(kws *ws.Websocket) {

		// Retrieve the user id from endpoint
		userId := kws.Params("id")

		// Add the connection to the list of the connected clients
		// The UUID is generated randomly and is the key that allow
		// ws to manage Emit/EmitTo/Broadcast
		clients[userId] = kws.UUID

		// Every websocket connection has an optional session key => value storage
		kws.SetAttribute("user_id", userId)

		//Broadcast to all the connected users the newcomer
		kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true)
		//Write welcome message
		kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)))
	}))

	go func() {
		for {
			time.Sleep(10 * time.Second)

			ws.Broadcast([]byte("string broadcast"))
		}
	}()

	log.Fatal(app.Listen(":3000"))
}
