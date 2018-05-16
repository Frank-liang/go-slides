package main

import "fmt"

func main() {
	gen := makeFibGen()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}

func makeFibGen() func() int {
	f1 := 0
	return func() int {
		f1++
		return f1
	}
}
