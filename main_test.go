package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_Dirigent(t *testing.T) {
	tests := []struct {
		name         string
		description  string
		route        string
		expectedCode int
	}{
		// Eerste test: Static page
		{
			name:         "Static",
			description:  "Get HTTP status: 200",
			route:        "/test-http",
			expectedCode: 200,
		},
	}

	app := fiber.New()

	app.Get("/test-http", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	for _, test := range tests {
		var (
			response *http.Response
		)

		request := httptest.NewRequest("GET", test.route, nil)

		response, _ = app.Test(request)

		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)
	}
}
