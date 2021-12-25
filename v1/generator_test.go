package v1

import (
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	sampler := func(x, y float64) float64 {
		if x == 0 && y == 0 {
			return 1
		} else if x == 0 && y == 18 {
			return 2
		} else if x == 18 && y == 18 {
			return 4
		} else if x == 18 && y == 0 {
			return 8
		}
		return 0
	}
	got := Generate(sampler)

	southWest := Vertex{0, 0, 9.3 + 1}
	northWest := Vertex{0, 18, 9.3 + 2}
	northEast := Vertex{18, 18, 9.3 + 4}
	southEast := Vertex{18, 0, 9.3 + 8}

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
			southWest,
			northWest,
			Vertex{0, 18, 0},
		},
		// back face
		{
			Vertex{0, 18, 0},
			northWest,
			northEast,
			Vertex{18, 18, 0},
		},
		// right face
		{
			Vertex{18, 18, 0},
			northEast,
			southEast,
			Vertex{18, 0, 0},
		},
		// front face
		{
			Vertex{18, 0, 0},
			southEast,
			southWest,
			Vertex{0, 0, 0},
		},
		// top face
		{
			southEast,
			northEast,
			northWest,
			southWest,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
}
