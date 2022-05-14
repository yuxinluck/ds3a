package main

import "fmt"

/*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	//pretect = 保护节点
	protect := &ListNode{
		Val:  0,
		Next: nil,
	}
    
	//last = 上一组结尾
	last := protect

	//分组，按组遍历
	for head != nil {
		//找到组尾部
		end := getEnd(head, k)
        if end == nil { break}
	    //nextHead 记录下一组头节点
		nextHead := end.Next
		//翻转
		reverseList(head,end)
		//end 变队头
		last.Next = end
		//head 变队尾
		head.Next = nextHead
		//更新每组前一组跟后一组之间的边
		last = head
		head = nextHead
	}
	return protect.Next
}

func getEnd(head *ListNode, k int) *ListNode {
	for head != nil {
		k--
		if k == 0 {
			break
		}
		head = head.Next
	}
	return head

}

//head到end 翻转
func reverseList(head *ListNode, end *ListNode)  {

	if head == end {return}
	last := head
	head = head.Next

	for head != end {
		n := head.Next
		head.Next = last
		last = head
		head = n
	}

	end.Next = last
}

func makeList(l []int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	for i := 0; i < len(l); i++ {
		node := &ListNode{
			Val:  l[i],
			Next: nil,
		}
		cur.Next = node
		cur = node
	}
	return head.Next
}

func main() {

	l := []int{1, 2, 3, 4, 5}

	head := makeList(l)

	head = reverseKGroup(head,2)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
