package main

import (
	"fmt"
	"strings"
)

func main() {
	// START OMIT
	s := "one two three one"
	var words []string
	words = strings.Fields(s)
	// END OMIT

	counter := make(map[string]int)

	for _, word := range words {
		counter[word]++
	}

	fmt.Println("counter is: ", counter)
}
