package main

import (
	"fmt"

	v1 "github.com/arctair/mare-orientale/v1"
)

func main() {
	polygons := v1.MakeFaces()
	polygons = v1.Triangulate(polygons)
	fmt.Print(v1.SerializeToSTL(polygons))
}
