package main

/*
https://leetcode.cn/problems/delete-node-in-a-bst/
450. 删除二叉搜索树中的节点
*/

/*
BST
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		next := root.Right
		//获取右边最小值
		for next.Left != nil {
			next = next.Left
		}
		//右边删除最小值，因为最小值孩子节点为空，可以直接删除
		root.Right = deleteNode(root.Right, next.Val)
		//root 位置不变，改变值为右边最小值
		root.Val = next.Val
		return root
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
