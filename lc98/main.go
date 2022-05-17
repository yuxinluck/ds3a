package main

import "math"

/*
98. 验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	var check func(root *TreeNode, rangeL int, rangeR int) bool

	check = func(root *TreeNode, rangeL, rangeR int) bool {
		if root == nil {
			return true
		}
		if root.Val < rangeL || root.Val > rangeR {
			return false
		}
		return check(root.Left, rangeL, root.Val-1) && check(root.Right, root.Val+1, rangeR)
	}

	return check(root, math.MinInt, math.MaxInt)
}
