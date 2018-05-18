package main

import (
	"bufio"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		os.Stdout.Write(append([]byte(line), '\n'))
	}
}
