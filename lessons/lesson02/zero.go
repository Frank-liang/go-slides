package main

import "fmt"

func main() {
	var (
		x int
		y float32
		z string
		p *int
	)
    
    //var num = 10
    //p = &num

	fmt.Printf("%v\n", x)
	//fmt.Printf("%v\n", &x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	//fmt.Printf("%v\n", *p)
}
