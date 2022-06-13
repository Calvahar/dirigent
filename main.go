package main

import (
	"math/rand"
	"time"

	"github.com/Proftaak-Semester-2/dirigent/src/configs"
	"github.com/Proftaak-Semester-2/dirigent/src/middleware"
	"github.com/Proftaak-Semester-2/dirigent/src/routes"
	"github.com/Proftaak-Semester-2/dirigent/src/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Zet een willekeurige seed die overal
	rand.Seed(time.Now().UnixNano())

	// Config voor de fiber app wordt geladen...
	config := configs.FiberConfig()

	// ...en wordt hier toegepast. De app wordt dan gebruikt om routes en middleware op toe te passen
	app := fiber.New(config)

	// Gebruik middleware van Fiber
	middleware.FiberMiddleware(app)

	// Maak de verschillende routes...
	routes.NormalRoutes(app)
	routes.NotFoundRoute(app)
	routes.Listeners()

	// ...en start de server dan op
	utils.StartServer(app)
}
