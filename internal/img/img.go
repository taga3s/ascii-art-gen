package img

import (
	"image"
	"image/png"
	"os"
)

// Load loads an image from a file.
func Load(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return img
}

// Resize resizes an image to a given width and height.
func Resize(img image.Image, magnification float64) (image.Rectangle, int, int) {
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

	return rect, w, h
}

// UnSync unsynchronizes an image file.
func UnSync(dest *image.RGBA) {
	tmp, err := os.Create("tmp.png")
	if err != nil {
		panic(err)
	}
	defer tmp.Close()
	png.Encode(tmp, dest)
	os.Remove("tmp.png")
}
