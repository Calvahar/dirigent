package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"

	"net/http"

	"log"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	// app.Get("/dashboard", monitor.New())

	app.Use("/ws", func(c *fiber.Ctx) error {
		// Source: Fiber Documentation
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// IoT device stuurt verzoek naar /ws
		// WebSocket verbinding wordt geopened.

		// Wij gebruiken de Headers van dat verzoek en roepen authReq aan

		// Antwoord van de response

		// Mag WEL een verbinding houden, laat de verbinding open staan

		// Mag NIET een verbinding houden, drop de verbinding

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:3000/test", nil)
		req.Header.Set("API_KEY", "idiiwadiwaj")
		req.Header.Set("USERNAME", "test")
		res, _ := client.Do(req)

		// fmt.Println(res.Request.Header)
		fmt.Println(res.Request.Header)

		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("paniek")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		// fmt.Println(c.Request().Header)
		fmt.Println(c.GetRespHeaders())
		c.Set("KEY", "STRING")
		return c.SendString("a")
	})

	log.Fatal(app.Listen(":3000"))
}

// func authReq(app fiber.App, API_KEY string, USERNAME string) bool {
// 	// localhost
// 	// token := gestuurde api key
//
// 	// Auth: API_KEY
// 	// USERNAME: MijnIOTDevice

// 	return true
// }
