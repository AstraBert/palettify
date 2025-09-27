package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"testing"
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
		{"/colors", "POST", "testfiles/docker.jpg", true},
		{"/colors", "POST", "testfiles/gopher.png", true},
		{"/colors", "GET", "testfiles/italy.jpg", false},
	}
	app := Setup()
	for _, tc := range testCases {
		if tc.route == "/colors" {
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
				body, err := io.ReadAll(res.Body)
				if err != nil {
					t.Errorf("Expecting no error when reading the response body, got %s", err.Error())
					continue
				} else {
					if strings.Contains(string(body), "An error occurred while extracting the palette colors, try again with a different image") {
						t.Error("Request did not return any colors even when it is supposed to")
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
