package main

/*
https://leetcode.cn/problems/successor-lcci/
面试题 04.06. 后继者

设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。

如果指定节点没有对应的“下一个”节点，则返回null。

BST
后继
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	return getNext(root, p.Val)

}

func getNext(root *TreeNode, val int) *TreeNode {
	var ans *TreeNode
	curr := root

	for curr != nil {
		if curr.Val == val {
			if curr.Right != nil {
				//获取右侧最小值
				ans = curr.Right
				for ans.Left != nil {
					ans = ans.Left
				}
				break
			}
		}

		if curr.Val > val {
			if ans == nil || ans.Val > curr.Val {
				ans = curr
			}
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return ans
}
