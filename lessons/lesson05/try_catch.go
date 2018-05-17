package main

import "fmt"

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main() {
	// 抛出一个panic的异常，然后在defer中通过recover捕获这个异常
	Try(func() {
		panic("exception message")
	}, func(e interface{}) {
		fmt.Println("caught exception: ", e)
	})
}
