package main

import (
	asciiArt "ascii-art-gen/internal/ascii_art"
	"ascii-art-gen/internal/img"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	DefaultMagnification = 1.0
)

type Inputs struct {
	path          string
	threshold     int
	magnification float64
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "threshold",
				Aliases: []string{"t"},
				Usage:   "the threshold for ASCII Art Generation",
			},
			&cli.Float64Flag{
				Name:    "magnification",
				Aliases: []string{"m"},
				Usage:   "the magnification factor for ASCII Art Generation",
			},
		},
		Action: func(c *cli.Context) error {
			err := runApp(c)
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runApp(c *cli.Context) error {
	inputs := Inputs{}
	if c.NArg() != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	inputs.path = c.Args().Get(0)

	// Load the image
	srcImg, err := img.Load(inputs.path)
	if err != nil {
		return err
	}

	if c.IsSet("magnification") {
		inputs.magnification = c.Float64("magnification")
	} else {
		inputs.magnification = DefaultMagnification
	}

	// Resize the image
	resizedImg := img.Resize(srcImg, inputs.magnification)

	if c.IsSet("threshold") {
		inputs.threshold = c.Int("threshold")
	} else {
		inputs.threshold = asciiArt.CalcOTSUThreshold(resizedImg, resizedImg.Bounds().Dy(), resizedImg.Bounds().Dx())
	}

	// Generate the ASCII Art
	output := asciiArt.Generate(resizedImg, inputs.threshold)

	// Print the ASCII Art
	fmt.Print(output)

	err = img.UnSync(resizedImg)
	if err != nil {
		return err
	}

	return nil
}
