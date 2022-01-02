package v1

import (
	"reflect"
	"testing"
)

func TestTriangulator(t *testing.T) {
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
