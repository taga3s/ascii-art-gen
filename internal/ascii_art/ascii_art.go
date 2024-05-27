package asciiart

import (
	"image"
	"image/color"
	"strings"
)

const char = "8"

// Generate generates an ASCII art from an image.
func Generate(dest image.Image, threshold int) string {
	srcBounds := dest.Bounds()

	asciiArt := strings.Builder{}
	asciiArt.WriteString("\n")

	for y := 0; y < srcBounds.Max.Y; y += 2 {
		line := strings.Builder{}

		for x := 0; x < srcBounds.Max.X; x++ {
			c := color.GrayModel.Convert(dest.At(x, y))
			gray, _ := c.(color.Gray)
			if gray.Y < uint8(threshold) {
				line.WriteString(" ")
			} else {
				line.WriteString(char)
			}
		}

		if strings.Contains(line.String(), char) {
			asciiArt.WriteString(line.String())
			asciiArt.WriteString("\n")
		}
	}

	return asciiArt.String()
}
