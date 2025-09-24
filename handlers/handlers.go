package handlers

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	imageprocessing "github.com/AstraBert/palettify/image_processing"
	"github.com/gofiber/fiber/v2"
)

func ExtractColorsImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server error: " + err.Error(), "colors": nil})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server error: " + err.Error(), "colors": nil})
	}
	defer src.Close()
	colors, err := imageprocessing.ProcessImage(src)
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server error: " + err.Error(), "colors": nil})
	}
	colorMap := map[int]map[string]uint8{}

	for i, c := range colors {
		rgba, _ := c.(color.RGBA)
		colorMap[i] = map[string]uint8{"R": rgba.R, "G": rgba.G, "B": rgba.B, "A": rgba.A}
	}

	return c.JSON(fiber.Map{"status": 200, "message": "Colors generated correctly", "colors": colorMap})
}
