package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("123456789")
	md5sum := md5.Sum(data)
	fmt.Printf("md5:%v, len:%d\n", md5sum, len(md5sum))
}
