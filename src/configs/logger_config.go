package configs

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig() logger.Config {

	// Return de config die gebruikt kan worden in ./middleware/fiber_middleware.go
	// Deze is nu leeg.
	return logger.Config{}
}

type Config struct {
	Connect    bool
	Disconnect bool
	Error      bool
}

func LogConfig() Config {
	Connect, _ := strconv.ParseBool(os.Getenv("CONNECT"))
	Disconnect, _ := strconv.ParseBool(os.Getenv("DISCONNECT"))
	Error, _ := strconv.ParseBool(os.Getenv("ERROR"))

	return Config{Connect, Disconnect, Error}
}
