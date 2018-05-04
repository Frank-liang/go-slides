package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("a.txt")
    //f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND | os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello\n")
	//f.Seek(1, os.SEEK_SET) 
    //f.WriteString("###")
	f.Close()
}
