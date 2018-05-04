package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "hello " +
		"world!" // concat

	var c1 byte
	c1 = s1[0]

	s2 := s1[0 : len(s1)-1] // slice

	i, _ := strconv.Atoi("-42") // strconv

	f, _ := strconv.ParseFloat("3.1415", 64)
	i64, _ := strconv.ParseInt("-42", 10, 64)

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", s1, c1, s2, i, f, i64)
}
