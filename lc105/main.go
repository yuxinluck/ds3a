package main

/*
105. 从前序与中序遍历序列构造二叉树
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	var root *TreeNode
	var build func(Pl int, Pr int, Il int, Ir int) *TreeNode

	build = func(Pl, Pr, Il, Ir int) *TreeNode {
		if Pl > Pr {
			return nil
		}
		var root *TreeNode = &TreeNode{}
		root.Val = preorder[Pl]

		//mid是root在inorder中的位置
		var mid int = Il
		for inorder[mid] != root.Val {
			mid += 1
		}
		//inorder: Inl 至 (mid-1) 左子树中序
		//inorder: (mid+1) 至 Ir  右子树中序
		//preorder: (Pl+1) 至 Pl + (mid-Il)  左子树先序
		//preorder: Pl+(mid-Il)+1 至 Pr 右子树先序
		root.Left = build(Pl+1, Pl+(mid-Il), Il, (mid - 1))
		root.Right = build(Pl+(mid-Il)+1, Pr, (mid + 1), Ir)
		return root
	}
	root = build(0, len(preorder)-1, 0, len(inorder)-1)
	return root
}
