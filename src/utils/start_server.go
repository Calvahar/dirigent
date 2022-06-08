package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	fiberConnURL := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("|> De server kon niet gestart worden! Reden: %v", err)
	}
}
