package v1

import (
	"reflect"
	"testing"
)

func TestTriangulator(t *testing.T) {
	triangulate := Triangulator()
	got := triangulate([][]Vertex{
		{
			{1, 1, -1},
			{1, 1, 1},
			{-1, -1, 1},
			{-1, -1, -1},
		},
	})
	want := [][]Vertex{
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
}
