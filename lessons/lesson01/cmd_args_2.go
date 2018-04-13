package main

import (
	"fmt"
	"os"
)

func main() {
	// multiple variable assignment
	s, sep := "", ""

	// range return <idx,val> pair
	// key is ignored with '_' for avoiding compile error in Go

	// slice s[m:n], 0<=m<=n<=len(s)

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
