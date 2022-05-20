package main

/*
236. 二叉树的最近公共祖先

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	//ans 只记录最深的
	var ans *TreeNode
	var dfs func(root *TreeNode) []bool

	dfs = func(root *TreeNode) []bool {
		res := make([]bool, 2)
		if root == nil {
			return res
		}

		Lres := dfs(root.Left)
		Rres := dfs(root.Right)

		// 公共祖先同时包含p,q, ans只记录最深层
		res[0] = Lres[0] || Rres[0] || root.Val == p.Val
		res[1] = Lres[1] || Rres[1] || root.Val == q.Val

		if res[0] && res[1] && ans == nil {
			ans = root
		}
		return res
	}

	dfs(root)
	return ans
}
