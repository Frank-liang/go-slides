package main

import "fmt"

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := &Point{1, 2}
	p.ScaleBy(2)

	p1 := Point{1, 2}
	p2 := &p1
	p2.ScaleBy(2)

	p3 := Point{1, 2}
	p3.ScaleBy(2) // 等价于 (&p3).ScaleBy(2)
}
