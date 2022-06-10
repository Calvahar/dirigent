package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	// Lees de Read Timeout uit de omgevingsvariabelen (.env) en zet het om naar een getal
	timeoutCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	appName := os.Getenv("SERVER_APP_NAME")

	// Return de config die gebruikt kan worden in ./middleware/fiber_middleware.go
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(timeoutCount),
		AppName:     appName,
	}
}
