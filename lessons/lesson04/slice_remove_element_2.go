package main

import "fmt"

func removeDuplicate(s []string) []string {
	for i := range s {
		if i < len(s)-2 && s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			return removeDuplicate(s)
		}
	}
	return s
}

func main() {
	s := []string{"tom", "tom", "tom", "jack", "tom"}
	fmt.Println(removeDuplicate(s))
}
