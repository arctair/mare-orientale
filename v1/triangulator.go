package v1

import (
	"fmt"
	"math"
)

func Triangulator(polygons [][]Vector3) (triangles [][]Vector3, err error) {
	for _, polygon := range polygons {
		moreTriangles, err := triangulateOne(polygon)
		if err != nil {
			return nil, err
		}
		triangles = append(triangles, moreTriangles...)
	}
	return
}

func triangulateOne(polygon []Vector3) (polygons [][]Vector3, err error) {
	normal := polygonNormal(polygon)
	var index int
	for len(polygon) > 3 {
		index, err = findEar(polygon, normal)
		if err != nil {
			return
		}
		polygons = append(polygons, []Vector3{
			polygon[index],
			polygon[index+1],
			polygon[(index+2)%len(polygon)],
		})
		polygon = append(polygon[:index+1], polygon[index+2:]...)
	}
	return append(polygons, []Vector3{
		polygon[0],
		polygon[1],
		polygon[2],
	}), nil
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

func findEar(polygon []Vector3, normal Vector3) (int, error) {
	for index := range polygon {
		v0, v1, v2 := polygon[index],
			polygon[(index+1)%len(polygon)],
			polygon[(index+2)%len(polygon)]
		if isConvex(normal, v0, v1, v2) && !containsAny(normal, v0, v1, v2, append(polygon[index:], polygon[:index]...)[3:]) {
			return index, nil
		}
	}
	return 0, fmt.Errorf("no ears left in %+v", polygon)
}

func isConvex(n0, v0, v1, v2 Vector3) bool {
	n1 := cross(v0, v1, v2)
	return sameWay(n0.X, n1.X) && sameWay(n0.Y, n1.Y) && sameWay(n0.Z, n1.Z)
}

func sameWay(a, b float64) bool {
	return signEquals(a, b) || math.Abs(a-b) < 0.00001
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

func signEquals(a, b float64) bool {
	if a < 0 && b > 0 {
		return false
	} else if b < 0 && a > 0 {
		return false
	}
	return true
}

func containsAny(n, v0, v1, v2 Vector3, vertices []Vector3) bool {
	for _, vertex := range vertices {
		if isConvex(n, v0, v1, vertex) && isConvex(n, v1, v2, vertex) && isConvex(n, v2, v0, vertex) {
			return true
		}
	}
	return false
}
