package main

import "fmt"

/*
200. 岛屿数量
https://leetcode.cn/problems/number-of-islands/
*/

func numIslands(grid [][]byte) int {
	ans := 0
	//方向数组,上下左右
	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}

	// 行列
	m := len(grid)
	n := len(grid[0])

	//访问标记
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// fmt.Println(visited)
	var dfs func(i, j int)
	dfs = func(i, j int) {
		fmt.Println(i, j)

		visited[i][j] = true

		for d := 0; d < 4; d++ {
			nx := i + dx[d]
			ny := j + dy[d]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			if grid[nx][ny] == '1' && !visited[nx][ny] {
				dfs(nx, ny)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}

	return ans
}
