package main

import "fmt"

func double(x int) (result int) {
	defer func() {
		result += x
	}()
	return x + x
}

func main() {
	fmt.Println(double(2))
}
