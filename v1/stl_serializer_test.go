package v1

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSTLSerializer(t *testing.T) {
	t.Run("empty solid", func(t *testing.T) {
		var b bytes.Buffer
		SerializeToSTL(&b, nil)
		got := b.String()
		want := "solid \nendsolid"
		if got != want {
			t.Fatalf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("one facet", func(t *testing.T) {
		var b bytes.Buffer
		SerializeToSTL(&b, [][]Vector3{
			{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
		})
		got := b.String()
		wantFacet := `  facet normal 0 0 0
    outer loop
      vertex 0 1 2
      vertex 3 4 5
      vertex 6 7 8
    endloop
  endfacet`
		want := fmt.Sprintf("solid \n%s\nendsolid", wantFacet)
		if got != want {
			t.Fatalf("got '%s' want '%s'", got, want)
		}
	})
}
