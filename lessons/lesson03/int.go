package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// INT OMIT
	var (
		v1 int
		v2 int32 // (-2^16) ~ (2^16-1)
		v3 int64
		v4 uint
		v5 uint32 // 0 ~ 2^32 -1
		v6 uint64
	)

	// END OMIT

	fmt.Printf("v1 = %d, size = %d\n", v1, unsafe.Sizeof(v1))
	fmt.Printf("v2 = %d, size = %d\n", v2, unsafe.Sizeof(v2))
	fmt.Printf("v3 = %d, size = %d\n", v3, unsafe.Sizeof(v3))
	fmt.Printf("v4 = %d, size = %d\n", v4, unsafe.Sizeof(v4))
	fmt.Printf("v5 = %d, size = %d\n", v5, unsafe.Sizeof(v5))
	fmt.Printf("v6 = %d, size = %d\n", v6, unsafe.Sizeof(v6))

}
