package v1

import (
	"reflect"
	"testing"
)

func TestFaces(t *testing.T) {
	got := MakeFaces()
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
}
