package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
    "strconv"
)

func main() {
    path := "/proc"
    f,_ := os.Open(path)
    infos,_ := f.Readdir(-1)
    for _, info := range infos {
        // find pid folder
        _, err := strconv.Atoi(info.Name())
        if err != nil {
            continue 
        }
        if info.IsDir() {
            fname := path + "/" + info.Name() + "/cmdline"
            buf, err := ioutil.ReadFile(fname)
            if err != nil {
                log.Fatal(err) 
            }
            fmt.Printf("%s\t%s\n", info.Name(), string(buf))
        }
    }

}
