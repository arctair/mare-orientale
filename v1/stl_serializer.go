package v1

import (
	"fmt"
	"strconv"
)

func SerializeToSTL(polygons [][]Vector3) string {
	output := "solid \n"
	for _, polygon := range polygons {
		output = writeFacet(output, polygon)
	}
	output += "endsolid"
	return output
}

func writeFacet(output string, polygon []Vector3) string {
	output += "  facet normal 0 0 0\n    outer loop\n"
	for _, vertex := range polygon {
		output = writeVertex(output, vertex)
	}
	output += "    endloop\n  endfacet\n"
	return output
}

func writeVertex(output string, vertex Vector3) string {
	x, y, z := strconv.FormatFloat(vertex.X, 'f', -1, 64),
		strconv.FormatFloat(vertex.Y, 'f', -1, 64),
		strconv.FormatFloat(vertex.Z, 'f', -1, 64)
	return output + fmt.Sprintf("      vertex %s %s %s\n", x, y, z)
}
