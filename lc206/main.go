package main

import "fmt"

/*
https://leetcode.cn/problems/reverse-linked-list/
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000


进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？


*/

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var pre *ListNode

	for head != nil {
        n := head.Next
        head.Next = pre
		pre = head
		head = n
	}

	return pre
}

func makeList( l []int) *ListNode {
    head := &ListNode{
		Val: 0,
		Next: nil,
	}
	cur := head
	for i:=0 ; i< len(l);i++ {
        node := &ListNode{
			Val: l[i],
			Next: nil,
		}
		cur.Next = node
		cur = node
	}

	return head.Next
}

func main() {

	l := []int{1,2,3,4,5}

	head := makeList(l)

	head = reverseList(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	
}
