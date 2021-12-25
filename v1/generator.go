package v1

func Generate() [][]Vertex {
	return [][]Vertex{
		// bottom face
		{
			Vertex{18, 18, 0},
			Vertex{18, 0, 0},
			Vertex{0, 0, 0},
			Vertex{0, 18, 0},
		},
		// left face
		{
			Vertex{0, 0, 0},
			Vertex{0, 0, 9.3},
			Vertex{0, 18, 9.3},
			Vertex{0, 18, 0},
		},
		// back face
		{
			Vertex{0, 18, 0},
			Vertex{0, 18, 9.3},
			Vertex{18, 18, 9.3},
			Vertex{18, 18, 0},
		},
		// right face
		{
			Vertex{18, 18, 0},
			Vertex{18, 18, 9.3},
			Vertex{18, 0, 9.3},
			Vertex{18, 0, 0},
		},
		// front face
		{
			Vertex{18, 0, 0},
			Vertex{18, 0, 9.3},
			Vertex{0, 0, 9.3},
			Vertex{0, 0, 0},
		},
		// top face
		{
			Vertex{18, 0, 9.3},
			Vertex{18, 18, 9.3},
			Vertex{0, 18, 9.3},
			Vertex{0, 0, 9.3},
		},
	}
}

type Vertex struct {
	X, Y, Z float64
}
