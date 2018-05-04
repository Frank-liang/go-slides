package main

import (
	"fmt"
)

func main() {
	s1 := "hello中国"
	// s1[0] = 'x'

	//a := []byte(s1)
	a := []rune(s1)
	for k, v := range a {
		fmt.Println(k, string(v))
	}

	a[5] = '美'
	s2 := string(a)

	fmt.Println(s1)
	fmt.Println(s2)
}
