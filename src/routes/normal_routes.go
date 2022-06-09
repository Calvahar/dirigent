package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/Proftaak-Semester-2/dirigent/src/middleware"
	ws "github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func NormalRoutes(a *fiber.App) {
	path, _ := os.Getwd()
	a.Static("/play", fmt.Sprintf("%s\\static\\piano", path))

	a.Static("/piano", fmt.Sprintf("%s\\static\\client", path))

	a.Get("/middleman", middleware.WSMiddleware, ws.New(func(kws *ws.Websocket) {

		clients[kws.UUID] = kws.UUID

		kws.SetAttribute("user_id", kws.UUID)
	}))

	a.Get("/broadcaster", middleware.WSMiddleware, websocket.New(func(c *websocket.Conn) {
		var (
			msg []byte
			err error
		)

		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println(err)
				break
			}

			ws.Broadcast(msg)
		}
	}))
}
