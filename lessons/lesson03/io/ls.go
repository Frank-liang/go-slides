package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
        flag := "-"
        if info.IsDir() {
           flag = "d" 
        }
		fmt.Printf("%v %dB %s\n", flag, info.Size(), info.Name())
	}
	f.Close()

}
