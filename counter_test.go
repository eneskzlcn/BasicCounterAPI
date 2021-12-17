package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert" // add Testify package
	"net/http/httptest"
	"testing"
)

func TestCounterApi(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	count := 0
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		value        int
	}{
		// First test case
		{
			description:  "get HTTP status 200, when successfully get counter value",
			route:        "/counter",
			expectedCode: 200,
			value:        count,
		},
		// Second test case
		{
			description:  "get HTTP status 200, when successfully get increased counter value",
			route:        "/increase",
			expectedCode: 200,
			value:        count,
		},
		//Third test case
		{
			description:  "get HTTP status 200, when successfully get decreased counter value",
			route:        "/decrease",
			expectedCode: 200,
			value:        count,
		},
		//Fourth test case
		{
			description:  "get HTTP status 200, when successfully get reseted counter value",
			route:        "/reset",
			expectedCode: 200,
			value:        count,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Create route with GET method for test
	app.Get("/counter", func(c *fiber.Ctx) error {
		// Return simple string as response
		return c.JSON(count)
	})
	app.Get("/increase", func(c *fiber.Ctx) error {
		// Return simple string as response
		count++
		return c.JSON(count)
	})
	app.Get("/decrease", func(c *fiber.Ctx) error {
		// Return simple string as response
		count--
		return c.JSON(count)
	})
	app.Get("/reset", func(c *fiber.Ctx) error {
		// Return simple string as response
		count = 0
		return c.JSON(count)
	})

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
