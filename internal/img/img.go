package img

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

// `Load` loads an image from a file.
func Load(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// `Resize` resizes an image to a given width and height.
func Resize(img image.Image, magnification float64) *image.RGBA {
	rect := img.Bounds()

	w, h := 0, 0
	arg := 180.0 * magnification

	if rect.Dx() > rect.Dy() {
		m := arg / float64(rect.Dx())
		w = int(arg)
		h = int(float64(rect.Dy()) * m)
	} else {
		m := arg / float64(rect.Dy())
		w = int(float64(rect.Dx()) * m)
		h = int(arg)
	}

	dest := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.CatmullRom.Scale(dest, dest.Bounds(), img, rect, draw.Over, nil)

	return dest
}

// `UnSync` unsynchronizes an image file.
func UnSync(dest image.Image) error {
	tmp, err := os.Create("tmp.png")
	if err != nil {
		return err
	}
	defer tmp.Close()

	err = png.Encode(tmp, dest)
	if err != nil {
		return err
	}
	os.Remove("tmp.png")

	return nil
}
