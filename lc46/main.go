package main

import "fmt"

/*
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

func permute(nums []int) [][]int {

	ans := make([][]int, 0)
	n := len(nums)

	//标记已经选的元素
	used := make([]bool, n)
	//组合
	a := make([]int, 0)
	//cnt 已经选cnt个数
	var rescur func(cnt int, ans *[][]int)

	rescur = func(cnt int, ans *[][]int) {
		if cnt == n {
			*ans = append(*ans, append([]int{}, a...))
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				a = append(a, nums[i])
				used[i] = true
				rescur(cnt+1, ans)
				//还原现场
				used[i] = false
				a = a[0 : len(a)-1]
			}
		}
	}

	rescur(0, &ans)
	return ans
}

func main() {

	s := permute([]int{1, 2, 3})
	fmt.Println(s)
}
