package v1

import "testing"

func serializeToSTL(polygons [][]Vertex) string {
	return "solid \nendsolid"
}

func TestSTLSerializer(t *testing.T) {
	got := serializeToSTL(nil)
	want := "solid \nendsolid"
	if got != want {
		t.Fatalf("got '%s' want '%s'", got, want)
	}
}
