package main

import (
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin

	buf := make([]byte, 1024)
	for {
		n, _ := f.Read(buf)
		os.Stdout.Write(buf[:n])
	}
}
