package main

/*
329. 矩阵中的最长递增路径
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-increasing-path-in-a-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//行
var m int

//列
var n int

//出边数组
var to [][]int

//出度
var deg []int

func longestIncreasingPath(matrix [][]int) int {

	ans := -1
	//行列数
	m = len(matrix)
	n = len(matrix[0])
	//方向数组
	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}

	//记录每个点深度
	depth := make([]int, m*n)
	//出边数组
	to = make([][]int, m*n)
	//入度
	deg = make([]int, m*n)

	//初始化出边数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 4; k++ {
				ni := i + dx[k]
				nj := j + dy[k]
				if valid(ni, nj) && matrix[ni][nj] > matrix[i][j] {
					addEdge(num(i, j), num(ni, nj))
				}
			}
		}
	}

	queue := make([]int, 0)

	for i := 0; i < m*n; i++ {
		if deg[i] == 0 {
			queue = append(queue, i)
			depth[i] = 1
		}
	}

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		for _, y := range to[x] {
			//当点y的入度减为零，它的深度也最大
			deg[y]--
			depth[y] = Max(depth[y], depth[x]+1)
			if deg[y] == 0 {
				queue = append(queue, y)
			}
		}
	}

	for i := 0; i < m*n; i++ {
		ans = Max(ans, depth[i])
	}
	return ans

}

func num(i, j int) int {
	return i*n + j
}

func valid(i, j int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func addEdge(u, v int) {
	to[u] = append(to[u], v)
	deg[v]++
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
