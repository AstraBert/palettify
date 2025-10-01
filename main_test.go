package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func TestRoutes(t *testing.T) {
	testCases := []struct {
		route         string
		method        string
		payload       string
		methodAllowed bool
	}{
		{"/", "GET", "", true},
		{"/", "POST", "", false},
		{"/html/colors", "POST", "testfiles/docker.jpg", true},
		{"/html/colors", "POST", "testfiles/gopher.png", true},
		{"/html/colors", "GET", "testfiles/italy.jpg", false},
		{"/json/colors", "POST", "testfiles/docker.jpg", true},
		{"/json/colors", "POST", "testfiles/gopher.png", true},
		{"/json/colors", "GET", "testfiles/italy.jpg", false},
	}
	app := Setup()
	for _, tc := range testCases {
		if strings.HasSuffix(tc.route, "/colors") {
			file, _ := os.Open(tc.payload)
			defer file.Close()

			var requestBody bytes.Buffer
			writer := multipart.NewWriter(&requestBody)

			fileWriter, _ := writer.CreateFormFile("image", tc.payload)

			io.Copy(fileWriter, file)

			writer.Close()
			req, _ := http.NewRequest(
				tc.method,
				tc.route,
				&requestBody,
			)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			res, _ := app.Test(req, -1)
			if res.StatusCode != 405 && !tc.methodAllowed {
				t.Errorf("Expecting the route to return a 'Method not allowed' (405) status code, got %d", res.StatusCode)
			}
			if res.StatusCode >= 400 && tc.methodAllowed {
				t.Errorf("Expecting the route to behave normally, got status code %d", res.StatusCode)
			}
			if res.StatusCode == 200 {
				if strings.HasPrefix(tc.route, "/html") {
					body, err := io.ReadAll(res.Body)
					if err != nil {
						t.Errorf("Expecting no error when reading the response body, got %s", err.Error())
						continue
					} else {
						if strings.Contains(string(body), "An error occurred while extracting the palette colors, try again with a different image") {
							t.Error("Request did not return any colors even when it is supposed to")
						}
					}
				} else {
					body, err := io.ReadAll(res.Body)
					if err != nil {
						t.Errorf("Expecting no error when reading the response body, got %s", err.Error())
						continue
					} else {
						if !strings.Contains(string(body), "Colors generated correctly") {
							t.Error("Request did not return any colors even when it is supposed to")
						}
					}
				}
			}
		} else if tc.route == "/" {
			req, _ := http.NewRequest(
				tc.method,
				tc.route,
				nil,
			)
			res, _ := app.Test(req, -1)
			if res.StatusCode != 405 && !tc.methodAllowed {
				t.Errorf("Expecting the route to return a 'Method not allowed' (405) status code, got %d", res.StatusCode)
				continue
			}
			if res.StatusCode >= 400 && tc.methodAllowed {
				t.Errorf("Expecting the route to behave normally, got status code %d", res.StatusCode)
				continue
			}
			if res.StatusCode == 200 {
				body, err := io.ReadAll(res.Body)
				if err != nil {
					t.Errorf("Expecting no error when reading the response body, got %s", err.Error())
					continue
				} else {
					if !strings.Contains(string(body), `<a href="/" class="font-comic text-3xl text-blue-400">Palettify</a>`) {
						t.Errorf("The request did not yield the expected HTML content, it instead returned %s", string(body))
					}
				}
			}
		}
	}
}

func TestLimiter(t *testing.T) {
	app := Setup()
	limiterConfig := limiter.Config{
		Max: 10,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Track limit per IP address
		},
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"message": "Service is currently unavailable due to server overload, retry soon!",
			})
		},
	}
	app.Get("/helloworld", limiter.New(limiterConfig), func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello world!"})
	})

	for i := range 15 {
		req, _ := http.NewRequest(
			"GET",
			"/helloworld",
			nil,
		)
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req, -1)

		if i < 10 && res.StatusCode > 400 {
			t.Errorf("Expecting no server-side errors, got status code %d", res.StatusCode)
		} else if i > 10 && res.StatusCode != 503 {
			t.Errorf("Expecting status code to be 503 for in-excess requests, got %d", res.StatusCode)
		}
	}
}
