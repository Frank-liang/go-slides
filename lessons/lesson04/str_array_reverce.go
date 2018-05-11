package main

import "fmt"
import "strings"

func reverse(slice_s []string) {
	var n = len(slice_s)
	for i := 0; i < n/2; i++ {
		slice_s[i], slice_s[n-1-i] = slice_s[n-1-i], slice_s[i]
	}
}

func main() {
	//slice_s := strings.Fields("one two three")
	str_arr := []string{"one", "two", "three"}
	slice_s := str_arr[:]
	reverse(slice_s)
	fmt.Println(strings.Join(slice_s, " "))
}
