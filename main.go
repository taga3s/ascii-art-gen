package main

import (
	asciiArt "ascii-art-gen/internal/ascii_art"
	"ascii-art-gen/internal/img"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/image/draw"
)

type Inputs struct {
	path          string
	threshold     int
	magnification float64
}

func main() {
	inputs := Inputs{}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "threshold",
				Aliases:     []string{"t"},
				Usage:       "threshold for ASCII art generation",
				Value:       128,
				Destination: &inputs.threshold,
			},
			&cli.Float64Flag{
				Name:        "magnification",
				Aliases:     []string{"m"},
				Usage:       "magnification factor for ASCII art generation",
				Value:       1,
				Destination: &inputs.magnification,
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return fmt.Errorf("invalid number of arguments")
			}
			inputs.path = c.Args().Get(0)

			run(inputs)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(inputs Inputs) {
	srcImg := img.Load(inputs.path)

	rct, w, h := img.Resize(srcImg, inputs.magnification)

	dest := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.CatmullRom.Scale(dest, dest.Bounds(), srcImg, rct, draw.Over, nil)

	output := asciiArt.Generate(dest, inputs.threshold)

	fmt.Print(output)

	img.UnSync(dest)
}
