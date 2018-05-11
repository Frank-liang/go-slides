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
	}
	ages1["b"] = 2

	delete(ages1, "b")

	fmt.Println(ages, ages1)
	//fmt.Println(ages == ages1)
}
