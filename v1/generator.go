package v1

func Generate(sampler func(vector2 Vector2) float64, numberOfCuts int) [][]Vector3 {
	topVertices := [][]Vector3{}
	subdivisionLength := float64(18) / float64(numberOfCuts+1)
	for x := 0; x <= numberOfCuts+1; x += 1 {
		xmm := float64(x) * subdivisionLength
		topVertices = append(topVertices, []Vector3{})
		for y := 0; y <= numberOfCuts+1; y += 1 {
			ymm := float64(y) * subdivisionLength
			topVertices[len(topVertices)-1] = append(
				topVertices[len(topVertices)-1],
				Vector3{xmm, ymm, 9.3 + sampler(Vector2{xmm, ymm})},
			)
		}
	}
	return append(
		[][]Vector3{
			bottomFace(),
			leftFace(topVertices),
			backFace(topVertices),
			rightFace(topVertices),
			frontFace(topVertices),
		},
		topFaces(topVertices)...,
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

func leftFace(topVertices [][]Vector3) []Vector3 {
	vertices := []Vector3{{0, 0, 0}}
	vertices = append(vertices, topVertices[0]...)
	vertices = append(vertices, Vector3{0, 18, 0})
	return vertices
}

func backFace(topVertices [][]Vector3) []Vector3 {
	vertices := []Vector3{{0, 18, 0}}
	for _, row := range topVertices {
		vertices = append(vertices, row[len(row)-1])
	}
	vertices = append(vertices, Vector3{18, 18, 0})
	return vertices
}

func rightFace(topVertices [][]Vector3) []Vector3 {
	vertices := []Vector3{{18, 18, 0}}
	lastRow := topVertices[len(topVertices)-1]
	for index := len(lastRow) - 1; index >= 0; index-- {
		vertices = append(vertices, lastRow[index])
	}
	vertices = append(vertices, Vector3{18, 0, 0})
	return vertices
}

func frontFace(topVertices [][]Vector3) []Vector3 {
	vertices := []Vector3{{18, 0, 0}}
	for index := len(topVertices) - 1; index >= 0; index-- {
		vertices = append(vertices, topVertices[index][0])
	}
	vertices = append(vertices, Vector3{0, 0, 0})
	return vertices
}

func topFaces(topVertices [][]Vector3) (faces [][]Vector3) {
	for y := len(topVertices[0]) - 1; y > 0; y-- {
		for x := 0; x < len(topVertices)-1; x++ {
			faces = append(faces, [][]Vector3{
				{
					topVertices[x][y],
					topVertices[x][y-1],
					topVertices[x+1][y-1],
				},
				{
					topVertices[x][y],
					topVertices[x+1][y-1],
					topVertices[x+1][y],
				},
			}...)
		}
	}
	return
}
