package main

import "fmt"

/*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

/*
递归
每个元素选与不选
*/

var chosen []int
var n int

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	n = len(nums)
	recur(nums, 0, &ans)
	return ans
}

//golang 注意切片的传递
//https://studygolang.com/articles/30995
func recur(nums []int, i int, ans *[][]int) {
	//递归边界
	if i == n {
		// fmt.Println(chosen,i)
		*ans = append(*ans, append([]int(nil), chosen...))
		// fmt.Printf("%p\n",ans)
		return
	}

	/*每层都是相似的逻辑，选或者不选*/
	//不选，直接进入下一个索引
	recur(nums, i+1, ans)
	//选了，然后进入下一个索引
	chosen = append(chosen, nums[i])
	recur(nums, i+1, ans)

	//还原现场
	chosen = chosen[:len(chosen)-1]
}

func main() {
	nums := []int{1, 2, 3}

	fmt.Println(subsets(nums))
}
