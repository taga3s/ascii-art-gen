package main

import (
	asciiArt "ascii-art-gen/internal/ascii_art"
	"ascii-art-gen/internal/img"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
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
	srcImg, err := img.Load(inputs.path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	resizedImg := img.Resize(srcImg, inputs.magnification)

	output := asciiArt.Generate(resizedImg, inputs.threshold)

	fmt.Print(output)

	err = img.UnSync(resizedImg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
