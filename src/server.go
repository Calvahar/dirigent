package server

import (
	"fmt"
	"log"
	"os"
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

func RunServer() {
	config := LogConfig{Connect: false, Disconnect: false, Error: true, Logger: true}
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

	app.Get("/ws/:id?", ws.New(func(kws *ws.Websocket) {

		// Sla het id dat de gebruiker meegeeft op in een variabele
		userId := kws.Params("id")

		// De UUID wordt willekeurig gegenereerd en kan gebruikt worden door iki
		clients[userId] = kws.UUID

		kws.SetAttribute("user_id", userId)

		// Broadcast naar alle gebruikers. true: alle gebruikers behalve nieuwe, false: iedereen
		// kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true)

		// Emit stuurt het bericht alleen naar de huidige WebSocket client
		// kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)))
	}))

	// Draai een aparte non-blocking functie
	go func() {
		for {
			time.Sleep(1 * time.Second)

			ws.Broadcast([]byte("string broadcast"))
		}
	}()

	// Luister op localhost
	listenTo := fmt.Sprintf("%s:%s", os.Getenv("SERVER_BIND"), os.Getenv("SERVER_PORT"))
	fmt.Println(listenTo)
	log.Fatal(app.Listen(listenTo))
}
