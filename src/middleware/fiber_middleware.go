package middleware

import (
	"github.com/Proftaak-Semester-2/dirigent/src/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/websocket/v2"
)

func FiberMiddleware(a *fiber.App) {
	lConfig := configs.LoggerConfig()
	hConfig := configs.HelmetConfig()

	a.Use(
		// Logger logt alle binnenkomende requests
		logger.New(lConfig),

		// Helmet is een middleware die beschermt tegen veel voorkomende kwetsbaarheden.
		helmet.New(hConfig),
	)

}

// Middleware die gebruikt kan worden voor endpoints die als WebSocket moeten werken.
func WSMiddleware(c *fiber.Ctx) error {
	// Als het request WebSocket Headers hefet, mag het door naar de volgende endpoint
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	// Anders returnt een 426, Upgrade Required
	// Dit is in JSON, zodat het voor de IoT apparaten makkelijker te lezen is dan een HTML-pagina
	return fiber.ErrUpgradeRequired
}
