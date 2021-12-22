package v1

import (
	"errors"
	"os"

	"golang.org/x/image/tiff"
)

func main(file string) (string, error) {
	reader, err := os.Open(file)
	if err != nil {
		return "", err
	}
	image, err := tiff.Decode(reader)
	if err != nil {
		return "", err
	}
	bounds := image.Bounds()
	aspectRatio := float64(bounds.Dy()) / float64(bounds.Dx())
	if aspectRatio != 94.2/284.7 {
		return "", ErrWrongAspectRatio
	}
	return "", nil
}

var ErrWrongAspectRatio = errors.New("aspect ratio must be 284.7 x 94.2")

type Quad struct {
	V0, V1, V2, V3 Vertex
}

type Vertex struct {
	X, Y, Z int
}
