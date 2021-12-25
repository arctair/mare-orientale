package v1

import (
	"errors"
	"image/color"
	"os"

	"golang.org/x/image/tiff"
)

func Sampler(file string) (func(x, y float64) uint16, error) {
	reader, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	image, err := tiff.Decode(reader)
	if err != nil {
		return nil, err
	}
	bounds := image.Bounds()
	width, height := float64(bounds.Dx()), float64(bounds.Dy())
	aspectRatio := height / width
	if aspectRatio != 94.2/284.7 {
		return nil, ErrWrongAspectRatio
	}

	return func(x, y float64) uint16 { return image.At(int(x*width/284.7), int(y*height/94.2)).(color.Gray16).Y }, nil
}

var ErrWrongAspectRatio = errors.New("aspect ratio must be 284.7 x 94.2")
