package main

import (
	"fmt"
)

func toUpper(s string) string {
	buf := []byte(s)
	for i := 0; i < len(buf); i++ {
        if(buf[i] >= 'a' && buf[i] <= 'z') {
		    buf[i] = buf[i] - 32
        }
	}
	return string(buf)
}

func main() {
	s1 := "hello中国"
	fmt.Println(toUpper(s1))
}
