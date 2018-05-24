package main

type Point struct {
	X, Y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Path) Distance() float64 {
	// TODO
}

func main() {
	p := Path{Point{-1, 0}, {0, 0}, {3, 4}}
	fmt.Println(p.Distance())
}
