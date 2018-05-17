package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// START OMIT
func read(f *os.File) (string, error) {
	var total []byte
	buf := make([]byte, 1024)
	for {
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		total = append(total, buf[:]...)
	}
	return string(total), nil
}

func main() {
	f, _ := os.Open("a.txt")
	s, err := read(f)
	if err != nil {
		log.Fatalf("read error:%v", err)
	}
	fmt.Println(s)
}

// END OMIT
