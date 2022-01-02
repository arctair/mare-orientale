package v1

import (
	"reflect"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	t.Run("zero cuts across length", func(t *testing.T) {
		sampler := func(vector2 Vector2) float64 {
			return map[Vector2]float64{
				{0, 0}:   1,
				{0, 16}:  4,
				{16, 0}:  8,
				{16, 16}: 32,
				{32, 0}:  64,
				{32, 16}: 256,
			}[vector2]
		}

		got := Generate(sampler, 0, 32, 16, 7.3)

		southWest := Vector3{0, 0, 7.3 + 1}
		northWest := Vector3{0, 16, 7.3 + 4}
		north := Vector3{16, 16, 7.3 + 32}
		northEast := Vector3{32, 16, 7.3 + 256}
		southEast := Vector3{32, 0, 7.3 + 64}
		south := Vector3{16, 0, 7.3 + 8}

		want := [][]Vector3{
			// bottom face
			{
				Vector3{32, 16, 0},
				Vector3{32, 0, 0},
				Vector3{0, 0, 0},
				Vector3{0, 16, 0},
			},
			// left face
			{
				Vector3{0, 0, 0},
				southWest,
				northWest,
				Vector3{0, 16, 0},
			},
			// back face
			{
				Vector3{0, 16, 0},
				northWest,
				north,
				northEast,
				Vector3{32, 16, 0},
			},
			// right face
			{
				Vector3{32, 16, 0},
				northEast,
				southEast,
				Vector3{32, 0, 0},
			},
			// front face
			{
				Vector3{32, 0, 0},
				southEast,
				south,
				southWest,
				Vector3{0, 0, 0},
			},
			// top west face
			{
				northWest,
				southWest,
				south,
			},
			{
				northWest,
				south,
				north,
			},
			// top east face
			{
				north,
				south,
				southEast,
			},
			{
				north,
				southEast,
				northEast,
			},
		}
		Equal(t, want, got)
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

		got := Generate(sampler, 1, 18, 18, 9.3)

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
