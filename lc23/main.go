package main

import (
	"container/heap"
	"fmt"
)

/*
23. 合并K个升序链表
*/

/*
二插堆
小顶堆
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//小顶堆
// heap.Interface
/*
要使用内置函数操作
heap.Pop(h) ，heap.Push(h,l), len(h)

不能直接使用 h.Pop(),  h.Push(l), h.len()
*/

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) }

func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val <= h[j].Val }

func (h ListNodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {

	h := &ListNodeHeap{}
	heap.Init(h)
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}

	}
	head := &ListNode{}
	tail := head

	fmt.Println(h.Len())
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		tail.Next = node
		tail = tail.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return head.Next
}
