package main

import "fmt"
import "flag"

func fibonacci(num int) int {
	if num < 2 {
		return 1
	}

	return fibonacci(num-1) + fibonacci(num-2)
}

var num = flag.Int("n", 0, "the number")

func main() {
    flag.Parse()
    var n = *num
    var nums = fibonacci(n)
	//for i := 0; i < n; i++ {
    //	nums := fibonacci(i)
    //  fmt.Println(nums)
	//}
    fmt.Println(nums)
}
