package v1

import (
	"reflect"
	"testing"
)

func TestTriangulator(t *testing.T) {
	triangulate := Triangulator()
	t.Run("convex polygon", func(t *testing.T) {
		got := triangulate([][]Vector3{
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
		got := triangulate([][]Vector3{
			{
				{-2, 0, 2},
				{1, 0, 2},
				{1, 0, 1},
				{2, 0, 1},
				{2, 0, -2},
			},
		})
		want := [][]Vector3{
			{
				{-2, 0, 2},
				{1, 0, 2},
				{1, 0, 1},
			},
			{
				{1, 0, 1},
				{2, 0, 1},
				{2, 0, -2},
			},
			{
				{-2, 0, 2},
				{1, 0, 1},
				{2, 0, -2},
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %+v want %+v", got, want)
		}
	})
}
