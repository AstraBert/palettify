package handlers

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	imageprocessing "github.com/AstraBert/palettify/image_processing"
	"github.com/AstraBert/palettify/templates"
	"github.com/gofiber/fiber/v2"
)

func ExtractColorsImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	c.Set("Content-Type", "text/html")
	if err != nil {
		comp := templates.Colors([]color.RGBA{})
		return comp.Render(c.Context(), c.Response().BodyWriter())
	}
	src, err := file.Open()
	if err != nil {
		comp := templates.Colors([]color.RGBA{})
		return comp.Render(c.Context(), c.Response().BodyWriter())
	}
	defer src.Close()
	colors, err := imageprocessing.ProcessImage(src)
	if err != nil {
		comp := templates.Colors([]color.RGBA{})
		return comp.Render(c.Context(), c.Response().BodyWriter())
	}
	colorsRGBA := []color.RGBA{}

	for _, c := range colors {
		rgba, _ := c.(color.RGBA)
		colorsRGBA = append(colorsRGBA, rgba)
	}

	comp := templates.Colors(colorsRGBA)
	return comp.Render(c.Context(), c.Response().BodyWriter())
}

func ExtractColorsJSON(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Server error: " + err.Error(), "colors": nil})
	}
	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Server error: " + err.Error(), "colors": nil})
	}
	defer src.Close()
	colors, err := imageprocessing.ProcessImage(src)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Server error: " + err.Error(), "colors": nil})
	}
	colorMap := map[int]map[string]uint8{}

	for i, c := range colors {
		rgba, _ := c.(color.RGBA)
		colorMap[i] = map[string]uint8{"R": rgba.R, "G": rgba.G, "B": rgba.B, "A": rgba.A}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Colors generated correctly", "colors": colorMap})
}

func HomeRoute(c *fiber.Ctx) error {
	home := templates.Home()
	c.Set("Content-Type", "text/html")
	return home.Render(c.Context(), c.Response().BodyWriter())
}
