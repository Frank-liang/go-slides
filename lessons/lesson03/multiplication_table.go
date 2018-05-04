package main

import (
	"fmt"
	"os"
)

func main() {
    f, _ := os.Create("t.txt")
	i, j := 1, 1
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-2d\t", i, j, i*j)
			fmt.Fprintf(f, "%d*%d=%-2d\t", i, j, i*j)
		}
		fmt.Println()
        f.WriteString("\n")
	}
    f.Close()
}
