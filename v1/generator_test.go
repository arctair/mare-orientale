package v1

import (
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	got := Generate()
	want := [][]Vertex{
		// bottom face
		{
			Vertex{18, 18, 0},
			Vertex{18, 0, 0},
			Vertex{0, 0, 0},
			Vertex{0, 18, 0},
		},
		// left face
		{
			Vertex{0, 0, 0},
			Vertex{0, 0, 9.3},
			Vertex{0, 18, 9.3},
			Vertex{0, 18, 0},
		},
		// back face
		{
			Vertex{0, 18, 0},
			Vertex{0, 18, 9.3},
			Vertex{18, 18, 9.3},
			Vertex{18, 18, 0},
		},
		// right face
		{
			Vertex{18, 18, 0},
			Vertex{18, 18, 9.3},
			Vertex{18, 0, 9.3},
			Vertex{18, 0, 0},
		},
		// front face
		{
			Vertex{18, 0, 0},
			Vertex{18, 0, 9.3},
			Vertex{0, 0, 9.3},
			Vertex{0, 0, 0},
		},
		// top face
		{
			Vertex{18, 0, 9.3},
			Vertex{18, 18, 9.3},
			Vertex{0, 18, 9.3},
			Vertex{0, 0, 9.3},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
}
