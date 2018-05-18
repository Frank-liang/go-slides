package main

import (
	"io/ioutil"
	"os"
)

func main() {
	var f *os.File
	f, _ = os.Open("a.txt")
	defer f.Close()
	buf, e := ioutil.ReadAll(f)
	if e != nil {
		panic(e)
	}
	os.Stdout.Write(buf[:])
}
