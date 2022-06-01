package main

import (
	"log"
	"time"
	"os"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var message string

func main() {
	// Nieuwe Fiber applicatie
	app := fiber.New(fiber.Config{
		AppName: "Dirigent Server",
	})

	// Gebruik logger om binnenkomende requests te loggen
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Redirect alles wat op de root binnenkomt naar /connect
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/connect")
	})

	// Van boven naar beneden; dit is de eerste /connect in de code.
	app.Use("/connect", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			// Ga door naar de volgende /connect
			return c.Next()
		}
		// Return een 426 status
		return fiber.ErrUpgradeRequired
	})
 
	app.Get("/connect", websocket.New(func(c *websocket.Conn) {	
		// Constante loop voor elke verbonden client
		for {
			// Om voor elke client synchroon te lopen, pak de huidige milliseconden;
			// Gelijk aan 000: precieze minuut
			ms := time.Now().Format(".000")
			if ms == ".000" {
				// Pak de huidige kleur uit color.txt
				msg, _ := os.ReadFile("colorOverwrite.txt")

				if len(msg) < 1 {
					msg, _ = os.ReadFile("color.txt")
				}



				// Schrijf deze als message type 1 (TextMessage) naar de verbonden clients
				c.WriteMessage(1, msg)
				// fmt.Println(c.)

				// Wacht 500 milliseconden, zodat de if statement maximaal 1x per seconde true is
				time.Sleep(500 * time.Millisecond)
			}
		}
	}))

	// Wordt gebruikt voor de DEMO client
	app.Static("/dep", "./static", fiber.Static{
		Browse: true,
	})
	app.Get("/demo", func(c *fiber.Ctx) error {
		return c.Redirect("http://dirigent.hanskazan.space/dep/client.html")
	})

	// Wordt gebruikt om de DEMO client makkelijk te downloaden
	app.Get("/client", func(c *fiber.Ctx) error {
		return c.Download("./download/client.html")
	})

	// Luister naar requests op poort 3000 (huidige domein: dirigent.hanskazan.space)
	log.Fatal(app.Listen(":3000"))
}