package v1

import (
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	t.Run("zero cuts", func(t *testing.T) {
		sampler := func(vector2 Vector2) float64 {
			return map[Vector2]float64{
				{0, 0}:   1,
				{0, 7}:   2,
				{0, 14}:  4,
				{7, 0}:   8,
				{7, 7}:   16,
				{7, 14}:  32,
				{14, 0}:  64,
				{14, 7}:  128,
				{14, 14}: 256,
			}[vector2]
		}

		got := Generate(sampler, 0, 14, 7.3)

		southWest := Vector3{0, 0, 7.3 + 1}
		northWest := Vector3{0, 14, 7.3 + 4}
		northEast := Vector3{14, 14, 7.3 + 256}
		southEast := Vector3{14, 0, 7.3 + 64}

		want := [][]Vector3{
			// bottom face
			{
				Vector3{14, 14, 0},
				Vector3{14, 0, 0},
				Vector3{0, 0, 0},
				Vector3{0, 14, 0},
			},
			// left face
			{
				Vector3{0, 0, 0},
				southWest,
				northWest,
				Vector3{0, 14, 0},
			},
			// back face
			{
				Vector3{0, 14, 0},
				northWest,
				northEast,
				Vector3{14, 14, 0},
			},
			// right face
			{
				Vector3{14, 14, 0},
				northEast,
				southEast,
				Vector3{14, 0, 0},
			},
			// front face
			{
				Vector3{14, 0, 0},
				southEast,
				southWest,
				Vector3{0, 0, 0},
			},
			// top face
			{
				northWest,
				southWest,
				southEast,
			},
			{
				northWest,
				southEast,
				northEast,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})

	t.Run("one cut", func(t *testing.T) {
		sampler := func(vector2 Vector2) float64 {
			return map[Vector2]float64{
				{0, 0}:   1,
				{0, 9}:   2,
				{0, 18}:  4,
				{9, 0}:   8,
				{9, 9}:   16,
				{9, 18}:  32,
				{18, 0}:  64,
				{18, 9}:  128,
				{18, 18}: 256,
			}[vector2]
		}

		got := Generate(sampler, 1, 18, 9.3)

		southWest := Vector3{0, 0, 9.3 + 1}
		west := Vector3{0, 9, 9.3 + 2}
		northWest := Vector3{0, 18, 9.3 + 4}
		north := Vector3{9, 18, 9.3 + 32}
		northEast := Vector3{18, 18, 9.3 + 256}
		east := Vector3{18, 9, 9.3 + 128}
		southEast := Vector3{18, 0, 9.3 + 64}
		south := Vector3{9, 0, 9.3 + 8}
		center := Vector3{9, 9, 9.3 + 16}

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
				west,
				northWest,
				Vector3{0, 18, 0},
			},
			// back face
			{
				Vector3{0, 18, 0},
				northWest,
				north,
				northEast,
				Vector3{18, 18, 0},
			},
			// right face
			{
				Vector3{18, 18, 0},
				northEast,
				east,
				southEast,
				Vector3{18, 0, 0},
			},
			// front face
			{
				Vector3{18, 0, 0},
				southEast,
				south,
				southWest,
				Vector3{0, 0, 0},
			},
			// top north west
			{
				northWest,
				west,
				center,
			},
			{
				northWest,
				center,
				north,
			},
			// top north east
			{
				north,
				center,
				east,
			},
			{
				north,
				east,
				northEast,
			},
			// top south west
			{
				west,
				southWest,
				south,
			},
			{
				west,
				south,
				center,
			},
			// top south east
			{
				center,
				south,
				southEast,
			},
			{
				center,
				southEast,
				east,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
}
