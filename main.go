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
	rand.Seed(time.Now().UnixNano())

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.NormalRoutes(app)
	routes.NotFoundRoute(app)
	routes.Listeners()

	utils.StartServer(app)
}
