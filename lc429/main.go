package main

/*
https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
429. N 叉树的层序遍历
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

*/

/*
使用队列，先进先出，做 BFS ，Breath First Search
*/

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {

	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	q := make([][]*Node, 0)
	q = append(q, []*Node{root})

	for len(q) > 0 {
		l := q[0]
		v := make([]int, 0)
		q = q[1:]

		if len(l) > 0 {
			for _, n := range l {
				v = append(v, n.Val)
			}
			ans = append(ans, v)

			node := make([]*Node, 0)

			for _, n := range l {
				node = append(node, n.Children...)
			}
			q = append(q, node)
		}
	}

	return ans

}
