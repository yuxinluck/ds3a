package main

import (
	"fmt"
	"strings"
)

/*
51. N 皇后
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
题解

二维数组上的点，不能同行，不能同列
对角线点: i-j ,i+j 唯一

DFS 实现
*/

func solveNQueens(n int) [][]string {

	//存储答案
	ans := make([][]int, 0)
	//标记已经使用的列
	used := make([]bool, n)
	//标记i-j,i+j
	usedPlus := make(map[int]bool)
	usedMinus := make(map[int]bool)
	//
	temp := make([]int, 0)

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			// fmt.Println(temp)
			//注意注意，go的坑，引用与值
			ans = append(ans, append([]int{}, temp...))
			return
		}

		for col := 0; col < n; col++ {
			if !used[col] && !usedPlus[row+col] && !usedMinus[row-col] {
				used[col] = true
				usedPlus[row+col] = true
				usedMinus[row-col] = true
				temp = append(temp, col)
				dfs(row + 1)
				//还原现场
				used[col] = false
				usedPlus[row+col] = false
				usedMinus[row-col] = false
				temp = temp[:len(temp)-1]
			}
		}
	}

	dfs(0)
	// fmt.Println(ans)
	//按照题意要求输出
	res := make([][]string, 0)
	for _, a := range ans {
		r := make([]string, 0)
		for _, d := range a {
			t := make([]string, n)
			for i := 0; i < n; i++ {
				t[i] = "."
			}
			t[d] = "Q"
			r = append(r, strings.Join(t, ""))
		}
		res = append(res, r)
	}
	// fmt.Println(res)
	return res
}

func main() {
	fmt.Println(solveNQueens(4))
}
