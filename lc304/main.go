package main

/*
304. 二维区域和检索 - 矩阵不可变
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/range-sum-query-2d-immutable
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	
		//m行
		m := len(matrix)
		//n列
		n := len(matrix[0])
	
		//前缀和
		sum := make([][]int, m+1)
		for i := 0; i <= m; i++ {
			sum[i] = make([]int, n+1)
		}
	
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
			}
		}

		return NumMatrix{sum: sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	row1++
	col1++
	row2++
	col2++

	return this.sum[row2][col2] - this.sum[row1-1][col2] - this.sum[row2][col1-1] + this.sum[row1-1][col1-1]

}
