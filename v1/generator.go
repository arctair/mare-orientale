package v1

func Generate(sampler func(x, y float64) float64) [][]Vector3 {
	grid := [][]Vector3{
		{Vector3{0, 0, 9.3 + sampler(0, 0)}, Vector3{0, 18, 9.3 + sampler(0, 18)}},
		{Vector3{18, 0, 9.3 + sampler(18, 0)}, Vector3{18, 18, 9.3 + sampler(18, 18)}},
	}
	return [][]Vector3{
		// bottom face
		{
			Vector3{18, 18, 0},
			Vector3{18, 0, 0},
			Vector3{0, 0, 0},
			Vector3{0, 18, 0},
		},
		// left face
		{
			Vector3{0, 0, 0},
			grid[0][0],
			grid[0][1],
			Vector3{0, 18, 0},
		},
		// back face
		{
			Vector3{0, 18, 0},
			grid[0][1],
			grid[1][1],
			Vector3{18, 18, 0},
		},
		// right face
		{
			Vector3{18, 18, 0},
			grid[1][1],
			grid[1][0],
			Vector3{18, 0, 0},
		},
		// front face
		{
			Vector3{18, 0, 0},
			grid[1][0],
			grid[0][0],
			Vector3{0, 0, 0},
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
