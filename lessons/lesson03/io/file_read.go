package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
    //bytes, _ := ioutil.ReadFile(filename)
	f, _ := os.Open(filename)
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
