package routes

import (
	"fmt"

	"github.com/Proftaak-Semester-2/dirigent/src/configs"
	ws "github.com/antoniodipinto/ikisocket"
)

var clients = make(map[string]string)

func Listeners() {

	config := configs.LogConfig()

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
}
