package main

import (
	"log"
	"os"
)

func main() {
	// append mode
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Failed to create a.txt")
	}
	f.WriteString("hello\n")
	f.Close()
}
