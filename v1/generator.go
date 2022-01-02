package v1

func Generate(sampler func(vector2 Vector2) float64, numberOfLengthCuts int, width, length, height float64) [][]Vector3 {
	topVertices := [][]Vector3{}

	numberOfLengthEdges := numberOfLengthCuts + 1
	numberOfWidthEdges := numberOfLengthEdges * int(width/length)
	subdivisionLength := length / float64(numberOfLengthEdges)
	subdivisionWidth := width / float64(numberOfWidthEdges)
	for x := 0; x <= numberOfWidthEdges; x += 1 {
		xmm := float64(x) * subdivisionWidth
		topVertices = append(topVertices, []Vector3{})
		for y := 0; y <= numberOfLengthEdges; y += 1 {
			ymm := float64(y) * subdivisionLength
			topVertices[len(topVertices)-1] = append(
				topVertices[len(topVertices)-1],
				Vector3{xmm, ymm, height + sampler(Vector2{xmm, ymm})},
			)
		}
	}
	return append(
		[][]Vector3{
			bottomFace(width, length),
			leftFace(length, topVertices),
			backFace(width, length, topVertices),
			rightFace(width, length, topVertices),
			frontFace(width, topVertices),
		},
		topFaces(topVertices)...,
	)
}

func bottomFace(width, length float64) []Vector3 {
	return []Vector3{
		{width, length, 0},
		{width, 0, 0},
		{0, 0, 0},
		{0, length, 0},
	}
}

func leftFace(length float64, topVertices [][]Vector3) []Vector3 {
	vertices := []Vector3{{0, 0, 0}}
	vertices = append(vertices, topVertices[0]...)
	vertices = append(vertices, Vector3{0, length, 0})
	return vertices
}

func backFace(width, length float64, topVertices [][]Vector3) []Vector3 {
	vertices := []Vector3{{0, length, 0}}
	for _, row := range topVertices {
		vertices = append(vertices, row[len(row)-1])
	}
	vertices = append(vertices, Vector3{width, length, 0})
	return vertices
}

func rightFace(width, length float64, topVertices [][]Vector3) []Vector3 {
	vertices := []Vector3{{width, length, 0}}
	lastRow := topVertices[len(topVertices)-1]
	for index := len(lastRow) - 1; index >= 0; index-- {
		vertices = append(vertices, lastRow[index])
	}
	vertices = append(vertices, Vector3{width, 0, 0})
	return vertices
}

func frontFace(width float64, topVertices [][]Vector3) []Vector3 {
	vertices := []Vector3{{width, 0, 0}}
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
