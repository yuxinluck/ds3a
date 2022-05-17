package main

import "fmt"

/*
77. 组合
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。
*/

/*
递归
每个元素选与不选
*/

func combine(n int, k int) [][]int {

	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	ans := make([][]int,0)
	chosen := make([]int,0)

	var recur func(i int, ans *[][]int) 
    recur = func(i int, ans *[][]int) {
        //减枝
        if len(chosen) > k || len(chosen)+(n-i) < k {return}
        
        if i == n || len(chosen) == k {
			if len(chosen) == k {
				*ans = append(*ans, append([]int{},chosen...))
			}
			return
		}
		recur(i+1,ans)
		chosen = append(chosen, nums[i])
		recur(i+1,ans)
		chosen = chosen[0:len(chosen)-1]
	}

	recur(0,&ans)
	return ans
}

func main() {

	fmt.Println(combine(4,2))
}
