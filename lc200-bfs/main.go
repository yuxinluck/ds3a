package main

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

	var bfs func(i, j int)
	bfs = func(i, j int) {
		queue := make([][]int, 0)
		queue = append(queue, []int{i, j})
		visited[i][j] = true

		for len(queue) > 0 {
			x := queue[0][0]
			y := queue[0][1]
			queue = queue[1:]

			for i := 0; i < 4; i++ {
				nx := x + dx[i]
				ny := y + dy[i]
				if nx < 0 || ny < 0 || nx >= m || ny >= n {
					continue
				}
				if grid[nx][ny] == '1' && !visited[nx][ny] {
					queue = append(queue, []int{nx, ny})
					visited[nx][ny] = true
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				ans++
				bfs(i, j)
			}
		}
	}

	return ans
}
