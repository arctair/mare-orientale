package main

import (
	"fmt"

	v1 "github.com/arctair/mare-orientale/v1"
)

func main() {
	polygons, err := v1.MakeFaces("v1/fixtures/right_aspect_ratio.tif")
	if err != nil {
		panic(err)
	}
	polygons = v1.Triangulate(polygons)
	fmt.Print(v1.SerializeToSTL(polygons))
}
