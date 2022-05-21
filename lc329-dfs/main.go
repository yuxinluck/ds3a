package main

/*
329. 矩阵中的最长递增路径
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-increasing-path-in-a-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestIncreasingPath(matrix [][]int) int {

	ans := -1
	m := len(matrix)
	n := len(matrix[0])

	//方向数组
	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}

	depth := make([][]int, 0)
	for i := 0; i < m; i++ {
		depth = append(depth, make([]int, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			depth[i][j] = -1
		}
	}

	// visited := make([][]bool, 0)
	// for i := 0; i < m; i++ {
	// 	visited = append(visited, make([]bool, n))
	// }

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// if visited[i][j] {
		// 	return
		// }
		if depth[i][j] != -1 {
			return
		}
		// ismax := true
		// for d := 0; d < 4; d++ {
		// 	nx := i + dx[d]
		// 	ny := j + dy[d]
		// 	if nx < 0 || ny < 0 || nx >= m || ny >= n {
		// 		continue
		// 	}
		// 	if matrix[nx][ny] > matrix[i][j] {
		// 		ismax = false
		// 	}
		// }
		// if ismax {
		// 	visited[i][j] = true
		// 	ans = Max(ans,depth[i][j])
		// 	return
		// }

		dep := 1
		for d := 0; d < 4; d++ {
			nx := i + dx[d]
			ny := j + dy[d]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			if matrix[nx][ny] > matrix[i][j] {
				if depth[nx][ny] != -1 {
					dep = Max(depth[nx][ny]+1, dep)
				}
				if depth[nx][ny] == -1 {
					dfs(nx, ny)
					dep = Max(depth[nx][ny]+1, dep)
				}
			}
		}
		depth[i][j] = dep
		ans = Max(dep, ans)
		// fmt.Println(ans, dep)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if depth[i][j] == -1 {
				dfs(i, j)
			}
		}
	}

	return ans
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
