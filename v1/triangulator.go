package v1

func Triangulator() func([][]Vector3) [][]Vector3 {
	return triangulate
}

func triangulate(polygons [][]Vector3) (triangles [][]Vector3) {
	for _, polygon := range polygons {
		triangles = append(triangles, triangulateOne(polygon)...)
	}
	return
}

func triangulateOne(polygon []Vector3) (polygons [][]Vector3) {
	normal := polygonNormal(polygon)
	for len(polygon) > 3 {
		index := findEar(polygon, normal)
		polygons = append(polygons, []Vector3{
			polygon[index],
			polygon[index+1],
			polygon[index+2],
		})
		polygon = append(polygon[:index+1], polygon[index+2:]...)
	}
	return append(polygons, []Vector3{
		polygon[0],
		polygon[1],
		polygon[2],
	})
}

func polygonNormal(polygon []Vector3) (normal Vector3) {
	for index := range polygon {
		vertexNormal := cross(
			polygon[index],
			polygon[(index+1)%len(polygon)],
			polygon[(index+2)%len(polygon)],
		)
		normal.X += vertexNormal.X
		normal.Y += vertexNormal.Y
		normal.Z += vertexNormal.Z
	}
	return
}

func findEar(polygon []Vector3, normal Vector3) int {
	for index := range polygon {
		vertexNormal := cross(
			polygon[index],
			polygon[(index+1)%len(polygon)],
			polygon[(index+2)%len(polygon)],
		)
		if hemispaceEquals(normal, vertexNormal) {
			return index
		}
	}
	panic("bad geometry")
}

func cross(v0, v1, v2 Vector3) Vector3 {
	a, b := Vector3{v1.X - v0.X, v1.Y - v0.Y, v1.Z - v0.Z},
		Vector3{v2.X - v1.X, v2.Y - v1.Y, v2.Z - v1.Z}
	return Vector3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func hemispaceEquals(n0, n1 Vector3) bool {
	return signEquals(n0.X, n1.X) && signEquals(n0.Y, n1.Y) && signEquals(n0.Z, n1.Z)
}

func signEquals(a, b float64) bool {
	if a < 0 && b > 0 {
		return false
	} else if b < 0 && a > 0 {
		return false
	}
	return true
}
