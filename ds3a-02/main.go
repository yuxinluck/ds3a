package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
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

func reverseList(head *ListNode,end *ListNode,t *ListNode)  {
	// pre := &ListNode{}
	
	cur := head
	pre := end.Next
    stop := end.Next

	for cur != stop {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	t.Next = end
}

// func reverseKGroup(head *ListNode, k int) *ListNode {
//     protect := &ListNode{}

//     end := head
//     start := head
//     t := protect

//     for end.Next != nil {
//         for i:=1;i<=k;i++ {
//            end = end.Next
//            if end == nil && i < k {
//                break
//            } 
//         }
//         reverseList(start,end,t)
//         t = end
//     }
// }



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

	t := &ListNode{}

	reverseList(n1,n5,t)

	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}

}