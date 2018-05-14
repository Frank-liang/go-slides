package main

import (
	"fmt"
	"strings"
)

func toupper(s string) string {
	mapFunc := func(r rune) rune {
		return r - ('a' - 'A')
	}
	return strings.Map(mapFunc, s)
}

func main() {
	fmt.Println(toupper("hello"))
}
