package main

import (
	"fmt"

	v1 "github.com/arctair/mare-orientale/v1"
)

func main() {
	faces, err := v1.MakeFaces("v1/fixtures/right_aspect_ratio.tif")
	if err != nil {
		panic(err)
	}
	fmt.Print(v1.SerializeToSTL(faces))
}
