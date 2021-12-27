package main

import (
	"fmt"
	"os"
	"strconv"

	v1 "github.com/arctair/mare-orientale/v1"
)

func main() {
	sampler, err := v1.Sampler(os.Args[1])
	check(err)
	numberOfCuts, err := strconv.Atoi(os.Args[2])
	check(err)
	polygons := v1.Generate(sampler, numberOfCuts)
	triangulator := v1.Triangulator()
	polygons = triangulator(polygons)
	fmt.Print(v1.SerializeToSTL(polygons))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
