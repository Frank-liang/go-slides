package main

import (
    "fmt"
    "sort"
)

type StringSlice []string
func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
    s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
    sort.Sort(StringSlice(s))
    fmt.Println(s)
}
