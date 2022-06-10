package routes

import (
	"fmt"
	"os"

	"github.com/Proftaak-Semester-2/dirigent/src/middleware"
	ws "github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
)

func NormalRoutes(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("dirigent.hanskazan.space/piano")
	})

	path, _ := os.Getwd()
	a.Static("/piano", fmt.Sprintf("%s\\static\\piano", path))

	a.Static("/connect", fmt.Sprintf("%s\\static\\client", path))

	a.Get("/middleman", middleware.WSMiddleware, ws.New(func(kws *ws.Websocket) {
		clients[kws.UUID] = kws.UUID

		kws.SetAttribute("user_id", kws.UUID)
	}))

	a.Get("/broadcaster", middleware.WSMiddleware, ws.New(func(kws *ws.Websocket) {
		if len(pianoPlayer) > 0 {
			kws.Emit([]byte("Someone is already using the piano!"))
			kws.Close()
		}

		pianoPlayer[kws.UUID] = kws.UUID
		kws.SetAttribute("user_id", kws.UUID)
	}))
}
