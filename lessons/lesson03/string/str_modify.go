package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "hello中国"
	// s1[0] = 'x'

	//a := []byte(s1)
	a := []rune(s1)
    fmt.Println(reflect.TypeOf(a))
	for k, v := range a {
		fmt.Println(k, string(v))
	}

	a[5] = '美'
	s2 := string(a)
}
