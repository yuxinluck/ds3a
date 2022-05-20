package main

/*
94. 二叉树的中序遍历
https://leetcode.cn/problems/binary-tree-inorder-traversal/
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	seq := make([]int, 0)
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		//左中右
		dfs(root.Left)
		seq = append(seq, root.Val)
		dfs(root.Right)
	}

	dfs(root)

	return seq
}
