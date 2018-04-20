package main

import (
    "fmt"
    "flag"
    "io/ioutil"
)

func printFile(name string) {
    buf, err := ioutil.ReadFile(name)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(buf))
}

var filePath = flag.String("f", " ", "file path")

func main() {
    flag.Parse()
    printFile(*filePath)
}
