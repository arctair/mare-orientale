package v1

func Generate(sampler func(vector2 Vector2) float64, numberOfCuts int) [][]Vector3 {
	grid := [][]Vector3{}
	subdivisionLength := float64(18) / float64(numberOfCuts+1)
	for x := 0.0; x <= 18.0; x += subdivisionLength {
		grid = append(grid, []Vector3{})
		for y := 0.0; y <= 18.0; y += subdivisionLength {
			grid[len(grid)-1] = append(
				grid[len(grid)-1],
				Vector3{x, y, 9.3 + sampler(Vector2{x, y})},
			)
		}
	}
	return append(
		[][]Vector3{
			bottomFace(),
			leftFace(grid),
			backFace(grid),
			rightFace(grid),
			frontFace(grid),
		},
		topFaces(grid)...,
	)
}

func bottomFace() []Vector3 {
	return []Vector3{
		{18, 18, 0},
		{18, 0, 0},
		{0, 0, 0},
		{0, 18, 0},
	}
}

func leftFace(samples [][]Vector3) []Vector3 {
	vertices := []Vector3{{0, 0, 0}}
	vertices = append(vertices, samples[0]...)
	vertices = append(vertices, Vector3{0, 18, 0})
	return vertices
}

func backFace(samples [][]Vector3) []Vector3 {
	vertices := []Vector3{{0, 18, 0}}
	for _, row := range samples {
		vertices = append(vertices, row[len(row)-1])
	}
	vertices = append(vertices, Vector3{18, 18, 0})
	return vertices
}

func rightFace(samples [][]Vector3) []Vector3 {
	vertices := []Vector3{{18, 18, 0}}
	lastRow := samples[len(samples)-1]
	for index := len(lastRow) - 1; index >= 0; index-- {
		vertices = append(vertices, lastRow[index])
	}
	vertices = append(vertices, Vector3{18, 0, 0})
	return vertices
}

func frontFace(samples [][]Vector3) []Vector3 {
	vertices := []Vector3{{18, 0, 0}}
	for index := len(samples) - 1; index >= 0; index-- {
		vertices = append(vertices, samples[index][0])
	}
	vertices = append(vertices, Vector3{0, 0, 0})
	return vertices
}

func topFaces(samples [][]Vector3) (faces [][]Vector3) {
	for y := len(samples[0]) - 1; y > 0; y-- {
		for x := 0; x < len(samples)-1; x++ {
			faces = append(faces, []Vector3{
				samples[x][y],
				samples[x][y-1],
				samples[x+1][y-1],
				samples[x+1][y],
			})
		}
	}
	return
}
