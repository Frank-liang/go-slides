package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	bytes, _ := ioutil.ReadFile(filename)
	fmt.Println(string(bytes))
}
