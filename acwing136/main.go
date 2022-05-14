package main

import (
	"fmt"
	"math"
	"sort"
)

/*

邻值查找
https://www.acwing.com/problem/content/description/138/

给定一个长度为 n 的序列 A，A 中的数各不相同。

对于 A 中的每一个数 Ai，求：

min1≤j<i|Ai−Aj|
以及令上式取到最小值的 j（记为 Pi）。若最小值点不唯一，则选择使 Aj 较小的那个。

输入格式
第一行输入整数 n，代表序列长度。

第二行输入 n 个整数A1…An,代表序列的具体数值，数值之间用空格隔开。

输出格式
输出共 n−1 行，每行输出两个整数，数值之间用空格隔开。

分别表示当 i 取 2∼n 时，对应的 min1≤j<i|Ai−Aj| 和 Pi 的值。

数据范围
n≤105,|Ai|≤109
输入样例：
3
1 5 3
输出样例：
4 1
2 1
*/

type Node struct {
	Val  int
	Idx  int
	Pre  *Node
	Next *Node
}

func newNode(val, idx int) *Node {
	return &Node{
		Val: val,
		Idx: idx,
	}
}

func addNode(p *Node, newNode *Node) {

	newNode.Pre = p
	newNode.Next = p.Next
	p.Next = newNode
}

func delNode(n *Node) {
	pre := n.Pre
	next := n.Next
	pre.Next = next
	next.Pre = pre
}

func main() {

	nums := []int{1, 5, 3}

	protectHead := newNode(0,-1e18)
	protectTail := newNode(0,1e18)
	addNode(protectHead,protectTail)

	//维护数组，按照索引顺序
	idxList := make([]*Node, len(nums))

	for i := 0; i < len(nums); i++ {
		idxList[i] = newNode(nums[i],i)
		// fmt.Println(idxList[i].Val,idxList[i].Idx)
	}

	//维护链表，按照值大小顺序
    valList := make([]*Node, len(nums))
	copy(valList,idxList)

	sort.Slice(valList, func(i,j int) bool {
		return valList[i].Val < valList[j].Val
	})

    p:= protectHead
	for i:=0;i<len(valList);i++ {
        addNode(p,valList[i])
		p = p.Next
	}

	// for i:=0 ; i< len(idxList);i++ {
	// 	fmt.Println(idxList[i].Val ,idxList[i].Idx)
	// }

	// for i:=0 ; i< len(valList);i++ {
	// 	fmt.Println(valList[i].Val , valList[i].Idx)
	// }


	// for p := protectHead.Next ; p != protectTail ; {
	// 	fmt.Println(p.Val , p.Idx)
	// 	p = p.Next
	// }

	for i:= len(idxList)-1; i > 0; i-- {
        
		if math.Abs(float64(idxList[i].Val - idxList[i].Pre.Val)) < math.Abs(float64(idxList[i].Val - idxList[i].Next.Val)) {
			fmt.Println(math.Abs(float64(idxList[i].Val - idxList[i].Pre.Val)), idxList[i].Pre.Idx)
		} else if math.Abs(float64(idxList[i].Val - idxList[i].Pre.Val)) == math.Abs(float64(idxList[i].Val - idxList[i].Next.Val)) {
			if idxList[i].Pre.Idx < idxList[i].Next.Idx {
				fmt.Println(math.Abs(float64(idxList[i].Val - idxList[i].Pre.Val)), idxList[i].Pre.Idx)
			} else {
				fmt.Println(math.Abs(float64(idxList[i].Val - idxList[i].Next.Val)), idxList[i].Next.Idx)
			}
		} else {
			fmt.Println(math.Abs(float64(idxList[i].Val - idxList[i].Next.Val)), idxList[i].Next.Idx)
		}
        
		// fmt.Println(idxList[i].Val)
		delNode(idxList[i])
	} 

}
