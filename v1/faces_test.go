package v1

import (
	"reflect"
	"testing"
)

func TestFaces(t *testing.T) {
	t.Run("wrong aspect ratio", func(t *testing.T) {
		_, err := faces("fixtures/wrong_aspect_ratio.tif")
		want := ErrWrongAspectRatio
		if err != want {
			t.Fatalf("got %s want %s", err, want)
		}
	})

	t.Run("right aspect ratio", func(t *testing.T) {
		got, err := faces("fixtures/right_aspect_ratio.tif")
		if err != nil {
			t.Fatal(err)
		}
		want := [][]Vertex{
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
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
}
