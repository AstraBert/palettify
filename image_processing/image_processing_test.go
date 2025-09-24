package imageprocessing

import (
	"image/color"
	"os"
	"testing"
)

func TestImageProcessing(t *testing.T) {
	testCases := []struct {
		filePath string
		expected []color.RGBA
	}{
		{"../testfiles/docker.jpg", []color.RGBA{{255, 255, 255, 255}, {36, 150, 237, 255}, {41, 87, 164, 255}}},
		{"../testfiles/germany.png", []color.RGBA{{217, 37, 39, 255}, {246, 205, 40, 255}, {20, 19, 19, 255}}},
		{"../testfiles/gopher.png", []color.RGBA{{106, 215, 229, 255}, {252, 242, 229, 255}, {3, 4, 4, 255}}},
		{"../testfiles/italy.jpg", []color.RGBA{{254, 255, 255, 255}, {205, 43, 55, 255}, {1, 146, 73, 255}}},
	}
	for _, tc := range testCases {
		imageFile, _ := os.Open(tc.filePath)
		defer imageFile.Close()
		inferredColors, err := ProcessImage(imageFile)
		if err != nil {
			t.Errorf("Expecting no error during the image processing, got %s", err.Error())
			continue
		}
		if len(tc.expected) != len(inferredColors) {
			t.Errorf("Expecting %d colors, got %d", len(tc.expected), len(inferredColors))
			continue
		}
		for i, c := range tc.expected {
			rgbaCol, isRGBA := inferredColors[i].(color.RGBA)
			if !isRGBA {
				t.Errorf("Expecting inferred colors to be of type color.RGBA, but color %d is not", i+1)
				continue
			}
			if c.R != rgbaCol.R || c.A != rgbaCol.A || c.B != rgbaCol.B || c.G != rgbaCol.G {
				t.Errorf("Expecting color %d to be %v, got %v", i+1, c, rgbaCol)
			}
		}
	}
}
