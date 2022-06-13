package routes

import (
	"fmt"

	"github.com/Proftaak-Semester-2/dirigent/src/configs"
	ws "github.com/antoniodipinto/ikisocket"
)

// Object van type <map> om verbindingen in op te slaan
var clients = make(map[string]string)
var pianoPlayer = make(map[string]string)

// Verschillende listeneres die werken op de ikisocket WebSocket.
func Listeners() {
	// Laadt de config en slaat het op in de config variabele
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
		delete(pianoPlayer, ep.Kws.GetStringAttribute("user_id"))

		// Als Disconnect in de config true is:
		if config.Disconnect {
			fmt.Printf("[Disconnect] - User: %s \n", ep.Kws.GetStringAttribute("user_id"))
		}
	})

	// Client's verbinding krijgt een error
	ws.On(ws.EventError, func(ep *ws.EventPayload) {
		// Als Error in de config true is:
		if config.Error {
			fmt.Printf("[Error] - User: %s \n", ep.Kws.GetStringAttribute("user_id"))
		}
	})

	ws.On(ws.EventMessage, func(ep *ws.EventPayload) {
		if len(pianoPlayer[ep.Kws.GetStringAttribute("user_id")]) > 0 {
			ws.Broadcast(ep.Data)
		}
	})
}
