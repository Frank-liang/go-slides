package main

import (
    "fmt"
    "flag"
)

func fib(n int) int {
    x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

var num = flag.Int("n", 0, "the number")

func main() {
    flag.Parse()
    var f = fib(*num)
    fmt.Printf("%d\n", f)
}
