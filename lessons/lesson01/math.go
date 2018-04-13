package main

import (
	"fmt"
	"math"
)

func main() {
	a := -2.0 //float64

	// godoc math
	// https://godoc.org/math

	fmt.Println(math.Abs(a)) // 2

	b := math.Pow(a, 2)
	fmt.Println(b) // 4

	fmt.Println(math.Sqrt(b)) // 2

}
