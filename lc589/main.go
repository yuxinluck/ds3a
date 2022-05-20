package main

/*
589. N 叉树的前序遍历
给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。

n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/n-ary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	seq := make([]int, 0)

	var dfs func(root *Node)

	dfs = func(root *Node) {
		if root == nil {
			return
		}
		seq = append(seq, root.Val)
		for _, v := range root.Children {
			dfs(v)
		}
	}

	dfs(root)
	return seq
}
