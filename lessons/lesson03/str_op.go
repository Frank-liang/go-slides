package main

import (
	"fmt"
)

func main() {
	// 相加
	s1 := "hello " + "world"
	fmt.Println(s1)

	// 取字符
	var c1 byte
	c1 = s1[0]
	fmt.Println(c1)
	fmt.Println(string(c1))
	fmt.Printf("%d %c\n", c1, c1)

	// 切片
	s2 := s1[0:len(s1)]
	fmt.Println(s2)
}
