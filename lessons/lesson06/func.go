package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{0, 0}
	q := Point{3, 4}
	fmt.Println(Distance(p, q)) // function call
	fmt.Println(p.Distance(q))
}
