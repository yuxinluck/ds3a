package main

import (
	"strconv"
	"strings"
)

/*
https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
297. 二叉树的序列化与反序列化
*/

/*
前序遍历

使用nil 补全叶子结点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	seq []string
}

func Constructor() Codec {
	return Codec{seq: []string{}}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			this.seq = append(this.seq, "NULL")
			return
		}

		val := root.Val
		this.seq = append(this.seq, strconv.Itoa(val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(this.seq, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.seq = strings.Split(data, ",")
	pos := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		val := this.seq[pos]
		pos++
		if val == "NULL" || pos >= len(this.seq) {
			return nil
		}
		v, _ := strconv.Atoi(val)
		root := &TreeNode{
			Val:   v,
			Left:  dfs(),
			Right: dfs(),
		}
		return root
	}

	root := dfs()
	return root
}
