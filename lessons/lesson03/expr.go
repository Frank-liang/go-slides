package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//for k,v := range os.Args {
	//    fmt.Println(k, v)
	//}

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run expr.go 1 + 2")
	}

	x, _ := strconv.Atoi(os.Args[1])
	op := os.Args[2]
	y, _ := strconv.Atoi(os.Args[3])

	//fmt.Println("operation is: ", op)

	switch op {
	case "+":
		fmt.Printf("x %s y = %d\n", op, x+y)
	case "-":
		fmt.Printf("x %s y = %d\n", op, x-y)
	case "*":
		fmt.Printf("x %s y = %d\n", op, x*y)
	case "/":
		fmt.Printf("x %s y = %d\n", op, x/y)
	}
}
