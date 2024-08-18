package asciiart

import (
	"image"
	"image/color"
	"math/rand"
	"strings"
)

const (
	validChars = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func selectRandomly(chars string) string {
	r := rand.Intn(len(chars))
	return string(chars[r])
}

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
				line.WriteString(selectRandomly(validChars))
			}
		}

		if len(strings.Fields(line.String())) != 0 {
			asciiArt.WriteString(line.String())
			asciiArt.WriteString("\n")
		}
	}

	return asciiArt.String()
}
