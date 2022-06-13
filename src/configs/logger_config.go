package configs

import (
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
	return Config{Connect: false, Disconnect: false, Error: false}
}
