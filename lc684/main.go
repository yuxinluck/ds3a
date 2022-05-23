package main

/*
684. 冗余连接
https://leetcode.cn/problems/redundant-connection/
*/

func findRedundantConnection(edges [][]int) []int {
	//点数
	var n int
	for _, v := range edges {
		x := v[0]
		y := v[1]
		n = Max(Max(x, y), n)
	}
	/*
		图的深度优先遍历模板
	*/
	//无向图，出边数组
	to := make([][]int, n+1)
	//记录是否已经访问
	visited := make([]bool, n+1)
	hasCycle := false
	
	var dfs func(x int, fa int)
	dfs = func(x, fa int) {
		visited[x] = true
		//出边数组访问x能到的边的方法
		for _, y := range to[x] {
			if y == fa {
				continue
			} else if !visited[y] {
				dfs(y, x)
			} else {
				hasCycle = true
			}
		}
	}

	for _, e := range edges {
		x := e[0]
		y := e[1]
		//一对无向边，就是两队单向边
		to[x] = append(to[x], y)
		to[y] = append(to[y], x)
		//加一条边，就判断一次
		hasCycle = false
		dfs(x, 0)
		if hasCycle {
			return e
		}
		//注意还原
		visited = make([]bool, n+1)
	}
	return []int{}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
