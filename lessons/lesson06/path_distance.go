package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func pathDistance(p Path) float64 {
	sum := 0.0
	n := len(p)
	if n == 0 || n == 1 {
		return 0
	}
	for i := 1; i < n; i++ {
		pre := p[i-1]
		cur := p[i]
		sum += cur.Distance(pre)
	}
	return sum
}

func (p Path) Distance() float64 {
	sum := 0.0
	n := len(p)
	if n == 0 || n == 1 {
		return 0
	}
	for i := 1; i < n; i++ {
		pre := p[i-1]
		cur := p[i]
		sum += cur.Distance(pre)
	}
	return sum
}

func main() {
	p := Path{Point{-1, 0}, {0, 0}, {3, 4}}
	fmt.Println(p.Distance())
	fmt.Println(pathDistance(p))
}
