package imageprocessing

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"

	color_extractor "github.com/marekm4/color-extractor"
)

func ProcessImage(src io.Reader) ([]color.Color, error) {
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, err
	}
	return color_extractor.ExtractColors(img), nil
}
