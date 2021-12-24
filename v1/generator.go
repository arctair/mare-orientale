package v1

func Generate() [][]Vertex {
	return [][]Vertex{
		// bottom face
		{
			Vertex{18 / 2, 18 / 2, -9.3 / 2},
			Vertex{18 / 2, -18 / 2, -9.3 / 2},
			Vertex{-18 / 2, -18 / 2, -9.3 / 2},
			Vertex{-18 / 2, 18 / 2, -9.3 / 2},
		},
		// left face
		{
			Vertex{-18 / 2, -18 / 2, -9.3 / 2},
			Vertex{-18 / 2, -18 / 2, 9.3 / 2},
			Vertex{-18 / 2, 18 / 2, 9.3 / 2},
			Vertex{-18 / 2, 18 / 2, -9.3 / 2},
		},
		// back face
		{
			Vertex{-18 / 2, 18 / 2, -9.3 / 2},
			Vertex{-18 / 2, 18 / 2, 9.3 / 2},
			Vertex{18 / 2, 18 / 2, 9.3 / 2},
			Vertex{18 / 2, 18 / 2, -9.3 / 2},
		},
		// right face
		{
			Vertex{18 / 2, 18 / 2, -9.3 / 2},
			Vertex{18 / 2, 18 / 2, 9.3 / 2},
			Vertex{18 / 2, -18 / 2, 9.3 / 2},
			Vertex{18 / 2, -18 / 2, -9.3 / 2},
		},
		// front face
		{
			Vertex{18 / 2, -18 / 2, -9.3 / 2},
			Vertex{18 / 2, -18 / 2, 9.3 / 2},
			Vertex{-18 / 2, -18 / 2, 9.3 / 2},
			Vertex{-18 / 2, -18 / 2, -9.3 / 2},
		},
		// top face
		{
			Vertex{18 / 2, -18 / 2, 9.3 / 2},
			Vertex{18 / 2, 18 / 2, 9.3 / 2},
			Vertex{-18 / 2, 18 / 2, 9.3 / 2},
			Vertex{-18 / 2, -18 / 2, 9.3 / 2},
		},
	}
}

type Vertex struct {
	X, Y, Z float64
}
