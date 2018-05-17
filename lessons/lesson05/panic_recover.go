package main

import "fmt"

func f() {
	fmt.Println("1")
	panic("exception message")
	fmt.Println("2")
}

// 输出什么?
func main() {
	defer func() {
		fmt.Println("3")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("4")
	}()

	f()
}
