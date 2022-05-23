package main

/*
https://leetcode.cn/problems/successor-lcci/
面试题 04.06. 后继者

设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。

如果指定节点没有对应的“下一个”节点，则返回null。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1. 查找p
2. 查找p.Right.Left
*/
//记录查询路径
// func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

// 	pos, seq := find(root, p)
// 	// fmt.Println(pos, seq)

// 	if pos == -1 {
// 		return nil
// 	}
// 	node := seq[pos]
// 	if node.Right != nil {
// 		return getMinBST(node.Right)
// 	} else {
// 		for i := pos - 1; i >= 0; i-- {
// 			//找到第一个大的就是 后继，返回
// 			if seq[i].Val > node.Val {
// 				return seq[i]
// 			}
// 		}
// 	}
// 	return nil
// }

//记录前一个比较大值
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	node, pre := find(root, p)
	// fmt.Println(pos, seq)

	if node == nil {
		return nil
	}
	if node.Right != nil {
		return getMinBST(node.Right)
	} else {
		return pre
	}

}

//获取BST最小值，即最左侧最小值 递归
// func getMinBST(root *TreeNode) *TreeNode {
// 	if root.Left == nil {
// 		return root
// 	}
// 	return getMinBST(root.Left)
// }
//获取BST最小值，即最左侧最小值 迭代
func getMinBST(root *TreeNode) *TreeNode {
	ans := root
	for root.Left != nil {
		ans = root.Left
		root = root.Left
	}
	return ans
}

//查询，并记录查询路径
// func find(root *TreeNode, p *TreeNode) (int, []*TreeNode) {

// 	var dfs func(root *TreeNode)
// 	pos := -1
// 	seq := make([]*TreeNode, 0)

// 	dfs = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		seq = append(seq, root)
// 		if root.Val == p.Val {
// 			pos = len(seq) - 1
// 		}
// 		dfs(root.Left)
// 		dfs(root.Right)
// 	}
// 	dfs(root)
// 	return pos, seq
// }

//查询，只记录前一个最大值
func find(root *TreeNode, p *TreeNode) (*TreeNode, *TreeNode) {

	var dfs func(root *TreeNode)
	var node *TreeNode
	var preLarger *TreeNode

	dfs = func(root *TreeNode) {
		if root == nil || node != nil {
			return
		}
		if root.Val > p.Val {
			preLarger = root
		}
		if root.Val == p.Val {
			node = root
			//找到了，好，不用继续查找了
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return node, preLarger
}
