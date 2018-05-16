package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func main() {
	m["a"] = 1
	fmt.Println(lookup("a"))
}
