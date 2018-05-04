package main

import (
	"fmt"
)

func toUpper(s string) string {
	buf := []byte(s)
    // 补全代码
	return string(buf)
}

func main() {
	s1 := "hello中国"
	fmt.Println(toUpper(s1))
}
