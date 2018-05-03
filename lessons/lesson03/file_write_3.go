package main

import (
    "os"
    "log"
)

func main()  {
    f,err := os.OpenFile("a.txt",os.O_CREATE | os.O_RDWR, 0644)
    if err != nil {
        log.Fatal(err)
    }
    f.WriteString("hello\n")
    f.Seek(1,os.SEEK_SET) //表示文件的其实位置，从第二个字符往后写入。
    f.WriteString("$$$")
    f.Close()
}
