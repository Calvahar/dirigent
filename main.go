package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	// app.Get("/dashboard", monitor.New())

	app.Use("/ws", func(c *fiber.Ctx) error {
		// Source: Fiber Documentation
		// Kijk of de HTTP request de `Upgrade: websocket` headers heeft.
		// Voeg de API_KEY en USERNAME headers toe aan Locals.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			c.Locals("header_KEY", c.Get("API_KEY"))
			c.Locals("header_USERNAME", c.Get("USERNAME"))

			// Ga naar de volgende /ws endpoint
			return c.Next()
		}

		// Als de connectie geen Upgrade header met WebSocket heeft, stuur een error terug.
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// Sla de API_KEY en Username op in een variable
		api_key := ""
		username := ""
		if c.Query("key") != "" && c.Query("username") != "" {
			api_key = c.Query("key")
			username = c.Query("username")
		} else if c.Locals("header_KEY") != nil && c.Locals("header_USERNAME") != nil {
			api_key = fmt.Sprintf("%s", c.Locals("header_KEY"))
			username = fmt.Sprintf("%s", c.Locals("header_USERNAME"))
		}

		client := &http.Client{}

		// url: De auth server.
		url := "http://localhost:3000/test"
		req, _ := http.NewRequest("GET", url, nil)

		// Stuurt de API key en Username mee naar de auth server.
		req.Header.Set("API_KEY", api_key)
		req.Header.Set("USERNAME", username)

		res, err := client.Do(req)

		if err != nil {
			fmt.Println(err)
		}

		// Hou de connectie open, wacht op antwoord van de auth server.
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		// @TODO: verander string in Allow || Deny
		if string(body) == "Ja dat klopt" {
			// c.Conn.Close()
			for {

				// err = c.WriteMessage(mt, msg)
				// if err != nil {
				// 	log.Println("write:", err)
				// 	break
				// }
			}
		}
	}))

	//Test authenticatie
	app.Get("/test", func(c *fiber.Ctx) error {
		api_key := c.Get("API_KEY") == "test"
		username := c.Get("USERNAME") == "test"

		if username != true || api_key != true {
			return c.SendString("Klopt niet")
		}
		return c.SendString("Ja dat klopt")
	})

	// Luister naar localhost op poort `3000`
	log.Fatal(app.Listen(":3000"))
}
