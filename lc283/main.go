package main

import "fmt"

/*
https://leetcode.cn/problems/move-zeroes/

移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意: 必须在不复制数组的情况下原地对数组进行操作。
*/

func moveZeroes(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}

	for ; n < len(nums); n++ {
		nums[n] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
