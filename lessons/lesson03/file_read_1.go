package main

import (
    "os"
    "bufio"
    "io"
    "fmt"
    "log"
)

func main()  {
    filename := os.Args[1]
    f,err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    r := bufio.NewReader(f) 
    for {
        line,err := r.ReadString('\n') 
        if err == io.EOF{ 
            break
        }
        fmt.Print(line)
    }
    f.Close()
}
