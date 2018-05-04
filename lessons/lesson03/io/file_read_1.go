package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	f.Close()
}
