package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// pre := &ListNode{}
	var pre *ListNode
	c := head

	for c != nil {
		tmp := c.Next
		c.Next = pre
		pre = c
		c = tmp
	}
	return pre
}

func main() {

	n5 := &ListNode{
		Val: 5,
	}

	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}

	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}

	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}

	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}

	t := reverseList(n1)

	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}

}