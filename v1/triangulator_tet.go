package v1

import (
	"reflect"
	"testing"
)

func Triangulate(polygons [][]Vertex) (triangles [][]Vertex) {
	for _, polygon := range polygons {
		triangles = append(triangles, triangulateOne(polygon)...)
	}
	return
}

func triangulateOne(polygon []Vertex) [][]Vertex {
	return [][]Vertex{
		{
			polygon[0],
			polygon[1],
			polygon[2],
		},
		{
			polygon[0],
			polygon[2],
			polygon[3],
		},
	}
}

func TestTriangulator(t *testing.T) {
	got := Triangulate([][]Vertex{
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
