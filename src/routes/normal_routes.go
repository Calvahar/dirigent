package routes

import (
	"time"

	"github.com/Proftaak-Semester-2/dirigent/src/controllers"
	"github.com/Proftaak-Semester-2/dirigent/src/middleware"
	ws "github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
)

func NormalRoutes(a *fiber.App) {
	a.Static("/piano", "./static/piano")

	a.Get("/piano", controllers.Piano)

	a.Get("/connect", middleware.WSMiddleware, ws.New(func(kws *ws.Websocket) {

		clients[kws.UUID] = kws.UUID

		kws.SetAttribute("user_id", kws.UUID)
	}))

	go func() {
		for {
			color := controllers.GenerateColor()
			time.Sleep(1 * time.Second)

			ws.Broadcast([]byte(string(color)))

		}
	}()
}
