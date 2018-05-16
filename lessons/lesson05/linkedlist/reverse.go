package main

func ReverseList(p *Node) *Node {
	var prev, next *Node = nil, p.Next
	for p != nil {
		next = p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return prev
}
