package util

import (
    "fmt"
)

func Reverse(s string) string {
    a := []rune(s)
    //n := len(a)

    for k, v := range a {
        fmt.Println(k, v)    
    }

    // 补全代码

    s2 := string(a)
    return s2
}
