package main

import "fmt"

func main() {
	// map[keyType]valueType
	ages := make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2

	// or

	ages1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println(ages, ages1)
}
