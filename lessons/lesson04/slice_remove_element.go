package main

import "fmt"

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	//s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(remove(s, 1))
}
