package main

import (
	"io"
	"os"
)

func main() {
	f := os.Stdin
	if len(os.Args) > 1 {
		f, _ = os.Open(os.Args[1])
		defer f.Close()
	}
	io.Copy(os.Stdout, f)
}
