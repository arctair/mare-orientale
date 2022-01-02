package v1

import (
	"os"
	"reflect"
	"testing"
)

func TestTriangulator(t *testing.T) {
	t.Run("float probs", func(t *testing.T) {
		sampler, err := Sampler("../data/heightmap.tif")
		if err != nil {
			t.Fatal(err)
		}
		polygons := Generate(sampler, 256, 118.3, 18, 9.3)
		polygons, err = Triangulator(polygons)
		if err != nil {
			t.Fatal(err)
		}
		SerializeToSTL(os.Stdout, polygons)
	})
	t.Run("convex polygon", func(t *testing.T) {
		got, _ := Triangulator([][]Vector3{
			{
				{1, 1, -1},
				{1, 1, 1},
				{-1, -1, 1},
				{-1, -1, -1},
			},
		})
		want := [][]Vector3{
			{
				{1, 1, -1},
				{1, 1, 1},
				{-1, -1, 1},
			},
			{
				{1, 1, -1},
				{-1, -1, 1},
				{-1, -1, -1},
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})

	t.Run("concave polygon", func(t *testing.T) {
		got, _ := Triangulator([][]Vector3{
			{
				{0, 0, 0},
				{3, 0, 0},
				{3, 0, 2},
				{2, 0, 1},
				{1, 0, 2},
			},
		})
		want := [][]Vector3{
			{
				{3, 0, 0},
				{3, 0, 2},
				{2, 0, 1},
			},
			{
				{0, 0, 0},
				{3, 0, 0},
				{2, 0, 1},
			},
			{
				{0, 0, 0},
				{2, 0, 1},
				{1, 0, 2},
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
}
