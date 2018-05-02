package main

import (
    "fmt"
)

func main() {
	str1 := "hello"
	doc := `
即使换行也不影响
可以验证一下
类似python的'''
`
    fmt.Println("str1 is: ", str1)
    fmt.Println("doc is: ", doc)
}
