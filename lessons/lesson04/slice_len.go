package main

import "fmt"

func main() {
	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	Slice_a := Array_a[2:5]
	fmt.Printf("len=%d cap=%d %v\n", len(Slice_a), cap(Slice_a), Slice_a)
}
