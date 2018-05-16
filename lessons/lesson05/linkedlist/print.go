package main

import "fmt"

var a = 10

func printList(p *Node) {
	for {
		if p == nil {
			fmt.Println("p is nil.")
			break
		}
		fmt.Println((*p).Val)
		p = p.Next
	}
}
