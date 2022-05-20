package main

import "fmt"

/*
1245. 树的直径

给你这棵「无向树」，请你测算并返回它的「直径」：这棵树上最长简单路径的 边数。

我们用一个由所有「边」组成的数组 edges 来表示一棵无向树，其中 edges[i] = [u, v] 表示节点 u 和 v 之间的双向边。

树上的节点都已经用 {0, 1, ..., edges.length} 中的数做了标记，每个节点上的标记都是独一无二的。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/tree-diameter
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
两次深度优先遍历
第一次从任意点出发，找到距离他最远的点p
第二次从p出发，找到据他最远的点q
连结p,q的路径就是树的直径
*/

func treeDiameter(edges [][]int) int {
	//n, 先计算有你个点
	n := 0
	for _, v := range edges {
		x := v[0]
		y := v[1]
		n = Max(n, Max(x, y))
	}
	n++
	//出边数组
	to := make([][]int, 0)
	for i := 0; i < n; i++ {
		to = append(to, []int{})
	}

	for _, v := range edges {
		to[v[0]] = append(to[v[0]], v[1])
		to[v[1]] = append(to[v[1]], v[0])
	}
	fmt.Println(to)

	findFarthest := func(start int) []int {
		//使用队列，广度优先遍历
		queue := make([]int, 0)
		queue = append(queue, start)
		
		//记录每个节点深度
		depth := make([]int, n)
		for i := 0; i < n; i++ {
			depth[i] = -1
		}
        depth[start] = 0

		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]
			for _, y := range to[x] {
				if depth[y] != -1 {
					continue
				} //已经走过，跳过
				depth[y] = depth[x] + 1
				queue = append(queue, y)
			}
		}

		ans := start
		for i := 0; i < n; i++ {
			if depth[i] > depth[ans] {
				ans = i
			}
		}
		return []int{ans, depth[ans]}
	}

	findFarthest(0)
	//第一个点
	p := findFarthest(0)
	fmt.Println(p)
	//第二个点
	q := findFarthest(p[0])
	fmt.Println(q)
	return q[1]

}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	//[[0,1],[1,2],[0,3],[3,4],[2,5],[3,6]]
	e := make([][]int, 6)
	e[0] = []int{0, 1}
	e[1] = []int{1, 2}
	e[2] = []int{0, 3}
	e[3] = []int{3, 4}
	e[4] = []int{2, 5}
	e[5] = []int{3, 6}

	treeDiameter(e)
}
