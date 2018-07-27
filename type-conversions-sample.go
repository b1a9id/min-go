package main

import (
	"math"
	"fmt"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z uint = uint(f)

	i := 42
	a := float64(i)
	b := uint(f)
	fmt.Println(x, y, z, i, a, b)
}
