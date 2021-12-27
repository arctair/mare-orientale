package v1

import (
	"fmt"
	"io"
	"strconv"
)

func SerializeToSTL(writer io.Writer, polygons [][]Vector3) {
	fmt.Fprint(writer, "solid \n")
	for _, polygon := range polygons {
		writeFacet(writer, polygon)
	}
	fmt.Fprint(writer, "endsolid")
}

func writeFacet(writer io.Writer, polygon []Vector3) {
	fmt.Fprint(writer, "  facet normal 0 0 0\n    outer loop\n")
	for _, vertex := range polygon {
		writeVertex(writer, vertex)
	}
	fmt.Fprintf(writer, "    endloop\n  endfacet\n")
}

func writeVertex(writer io.Writer, vertex Vector3) {
	x, y, z := strconv.FormatFloat(vertex.X, 'f', -1, 64),
		strconv.FormatFloat(vertex.Y, 'f', -1, 64),
		strconv.FormatFloat(vertex.Z, 'f', -1, 64)
	fmt.Fprintf(writer, "      vertex %s %s %s\n", x, y, z)
}
