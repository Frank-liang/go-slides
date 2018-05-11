package main

import "fmt"

func main() {
	// START OMIT
	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	Slice_a := Array_a[2:5]
	fmt.Printf("len=%d cap=%d %v\n", len(Slice_a), cap(Slice_a), Slice_a)

	Slice_b := Slice_a[0:8]
	fmt.Printf("len=%d cap=%d %v\n", len(Slice_b), cap(Slice_b), Slice_b)
	// END OMIT
}
