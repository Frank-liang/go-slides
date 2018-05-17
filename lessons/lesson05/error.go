package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// START OMIT
func read(f *os.File) (string, error) {
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal("e:", err)
	}

	content, err := read(f)
	if err != nil {
		log.Fatal("e:", err)
	}
	fmt.Println(content)
}

// END OMIT
