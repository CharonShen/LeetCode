package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	var cur *ListNode
	fmt.Println(head.Val, head.Next, cur.Val, cur.Next)
}
