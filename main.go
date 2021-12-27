package main

import (
	"fmt"
	"os"

	v1 "github.com/arctair/mare-orientale/v1"
)

func main() {
	sampler, err := v1.Sampler(os.Args[1])
	if err != nil {
		panic(err)
	}
	polygons := v1.Generate(sampler)
	triangulator := v1.Triangulator()
	polygons = triangulator(polygons)
	fmt.Print(v1.SerializeToSTL(polygons))
}
