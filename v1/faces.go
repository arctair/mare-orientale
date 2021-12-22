package v1

import (
	"errors"
	"os"

	"golang.org/x/image/tiff"
)

func faces(file string) ([][]Vertex, error) {
	reader, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	image, err := tiff.Decode(reader)
	if err != nil {
		return nil, err
	}
	bounds := image.Bounds()
	aspectRatio := float64(bounds.Dy()) / float64(bounds.Dx())
	if aspectRatio != 94.2/284.7 {
		return nil, ErrWrongAspectRatio
	}
	return [][]Vertex{
		// bottom face
		{
			Vertex{18 / 2, 18 / 2, -9.3 / 2},
			Vertex{18 / 2, -18 / 2, -9.3 / 2},
			Vertex{-18 / 2, -18 / 2, -9.3 / 2},
			Vertex{-18 / 2, 18 / 2, -9.3 / 2},
		},
		// left face
		{
			Vertex{-18 / 2, -18 / 2, -9.3 / 2},
			Vertex{-18 / 2, -18 / 2, 9.3 / 2},
			Vertex{-18 / 2, 18 / 2, 9.3 / 2},
			Vertex{-18 / 2, 18 / 2, -9.3 / 2},
		},
		// back face
		{
			Vertex{-18 / 2, 18 / 2, -9.3 / 2},
			Vertex{-18 / 2, 18 / 2, 9.3 / 2},
			Vertex{18 / 2, 18 / 2, 9.3 / 2},
			Vertex{18 / 2, 18 / 2, -9.3 / 2},
		},
		// right face
		{
			Vertex{18 / 2, 18 / 2, -9.3 / 2},
			Vertex{18 / 2, 18 / 2, 9.3 / 2},
			Vertex{18 / 2, -18 / 2, 9.3 / 2},
			Vertex{18 / 2, -18 / 2, -9.3 / 2},
		},
		// front face
		{
			Vertex{18 / 2, -18 / 2, -9.3 / 2},
			Vertex{18 / 2, -18 / 2, 9.3 / 2},
			Vertex{-18 / 2, -18 / 2, 9.3 / 2},
			Vertex{-18 / 2, -18 / 2, -9.3 / 2},
		},
		// top face
		{
			Vertex{18 / 2, -18 / 2, 9.3 / 2},
			Vertex{18 / 2, 18 / 2, 9.3 / 2},
			Vertex{-18 / 2, 18 / 2, 9.3 / 2},
			Vertex{-18 / 2, -18 / 2, 9.3 / 2},
		},
	}, nil
}

var ErrWrongAspectRatio = errors.New("aspect ratio must be 284.7 x 94.2")

type Vertex struct {
	X, Y, Z float64
}
