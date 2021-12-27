package v1

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
