package main

import "fmt"

func min_max(args ...int) (min, max int) {
	min, max = args[0], args[0]
	for i := 0; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
		if args[i] > max {
			max = args[i]
		}
	}
	return
}

func main() {
	fmt.Println(min_max(1, 2, 3))
}
