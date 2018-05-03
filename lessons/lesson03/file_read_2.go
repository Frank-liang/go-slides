package main

import (
    "os"
    "io/ioutil"
    "fmt"
)

func main()  {
    filename := os.Args[1]
    bytes, _:= ioutil.ReadFile(filename)
    fmt.Println(string(bytes))    
}
