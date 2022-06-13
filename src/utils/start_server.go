package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	// Maak een URL met behulp van .env ; uiteindelijk komt er iets van 0.0.0.0:3000
	fiberConnURL := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	// Probeer te luisteren op de URL die gemaakt is
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("|> De server kon niet gestart worden! Reden: %v", err)
	}
}
