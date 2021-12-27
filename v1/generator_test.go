package v1

import (
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	sampler := func(x, y float64) float64 {
		return map[Vector2]float64{
			{0, 0}:   1,
			{0, 18}:  2,
			{18, 18}: 4,
			{18, 0}:  8,
		}[Vector2{x, y}]
	}
	got := Generate(sampler)

	southWest := Vector3{0, 0, 9.3 + 1}
	northWest := Vector3{0, 18, 9.3 + 2}
	northEast := Vector3{18, 18, 9.3 + 4}
	southEast := Vector3{18, 0, 9.3 + 8}

	want := [][]Vector3{
		// bottom face
		{
			Vector3{18, 18, 0},
			Vector3{18, 0, 0},
			Vector3{0, 0, 0},
			Vector3{0, 18, 0},
		},
		// left face
		{
			Vector3{0, 0, 0},
			southWest,
			northWest,
			Vector3{0, 18, 0},
		},
		// back face
		{
			Vector3{0, 18, 0},
			northWest,
			northEast,
			Vector3{18, 18, 0},
		},
		// right face
		{
			Vector3{18, 18, 0},
			northEast,
			southEast,
			Vector3{18, 0, 0},
		},
		// front face
		{
			Vector3{18, 0, 0},
			southEast,
			southWest,
			Vector3{0, 0, 0},
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
