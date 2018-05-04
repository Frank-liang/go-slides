package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	fmt.Println("x: ", x)
	fmt.Print("guess a number 1-10:")
	var n int
	fmt.Scanf("%d", &n)

	// 补全代码，如果n==x 输出"正确"
	// 如果n>x输出"过大"
	// 如果n<x输出"过小"
}
