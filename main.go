package main

import (
	"os"
	"strconv"

	v1 "github.com/arctair/mare-orientale/v1"
)

func main() {
	sampler, err := v1.Sampler(os.Args[1])
	check(err)
	numberOfCuts, err := strconv.Atoi(os.Args[2])
	check(err)
	polygons := v1.Generate(sampler, numberOfCuts, 18, 18, 9.3)
	triangulator := v1.Triangulator()
	polygons = triangulator(polygons)
	v1.SerializeToSTL(os.Stdout, polygons)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
