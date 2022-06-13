package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/Proftaak-Semester-2/dirigent/src/controllers"
	"github.com/Proftaak-Semester-2/dirigent/src/middleware"
	ws "github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
)

func NormalRoutes(a *fiber.App) {
	// Redirect alles wat op de root binnenkomt naar dirigent.hanskazan.space/piano
	a.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("dirigent.hanskazan.space/piano")
	})

	// Statische routes die de bestanden uit de absolute paden /static/piano en /static/client serven
	path, _ := os.Getwd()

	a.Static("/piano", fmt.Sprintf("%s/static/piano", path))
	a.Static("/connect", fmt.Sprintf("%s/static/client", path))
	a.Static("/demo", fmt.Sprintf("%s/static/demo", path))

	// De middleman endpoint gebruikt de WebSocket middleware en de ikisocket package
	/* De client verbindt met /middleman, waarna de client wordt toegevoegd in een array
	 * van alle verbonden clients. Ook wordt er een attribute op de huidige connectie gezet.
	 */
	a.Get("/middleman", middleware.WSMiddleware, ws.New(func(kws *ws.Websocket) {
		clients[kws.UUID] = kws.UUID

		kws.SetAttribute("user_id", kws.UUID)
	}))

	// De broadcaster endpoint gebruikt de WebSocket middleware en de ikisocket package
	/* De piano speler verbindt met /broadcaster. Als er al iemand piano speelt (een verbonden client in pianoPlayer)
	 * dan wordt er een bericht gestuurd en wordt de verbinding gesloten. Zo niet, dan wordt de huidige speler
	 * toegevoegd aan de lijst van pianoPlayers en wordt er een attribute toegevoegd aan de huidige connectie.
	 */
	a.Get("/broadcaster", middleware.WSMiddleware, ws.New(func(kws *ws.Websocket) {
		if len(pianoPlayer) > 0 {
			kws.Emit([]byte("Iemand is de piano al aan het gebruiken!"))
			kws.Close()
		}

		pianoPlayer[kws.UUID] = kws.UUID
		kws.SetAttribute("user_id", kws.UUID)
	}))

	a.Get("/demo", middleware.WSMiddleware, ws.New(func(kws *ws.Websocket) {
		for {
			kws.Emit([]byte(controllers.GenerateColor()))
			time.Sleep(2 * time.Second)
		}
	}))
}
