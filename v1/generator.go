package v1

func Generate(sampler func(x, y float64) float64) [][]Vertex {
	grid := [][]Vertex{
		{Vertex{0, 0, 9.3 + sampler(0, 0)}, Vertex{0, 18, 9.3 + sampler(0, 18)}},
		{Vertex{18, 0, 9.3 + sampler(18, 0)}, Vertex{18, 18, 9.3 + sampler(18, 18)}},
	}
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
			grid[0][0],
			grid[0][1],
			Vertex{0, 18, 0},
		},
		// back face
		{
			Vertex{0, 18, 0},
			grid[0][1],
			grid[1][1],
			Vertex{18, 18, 0},
		},
		// right face
		{
			Vertex{18, 18, 0},
			grid[1][1],
			grid[1][0],
			Vertex{18, 0, 0},
		},
		// front face
		{
			Vertex{18, 0, 0},
			grid[1][0],
			grid[0][0],
			Vertex{0, 0, 0},
		},
		// top face
		{
			grid[1][0],
			grid[1][1],
			grid[0][1],
			grid[0][0],
		},
	}
}

type Vertex struct {
	X, Y, Z float64
}
