package main

/*
589. N 叉树的前序遍历
给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。

n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/n-ary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
前序，中序，后续 都是深度优先，使用递归或迭代

本题使用迭代

使用 栈 ，后进先出，做dfs
*/
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	seq := make([]int, 0)
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if node != nil {
			seq = append(seq, node.Val)
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
		}
	}
	return seq
}
