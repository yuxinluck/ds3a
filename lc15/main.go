package main

import "sort"

/*
15. 三数之和
https://leetcode.cn/problems/3sum/
*/

/*
题解：
双指针
将三数之和分解为两数之和
*/

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		//两数之和去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		s := twoSum(nums, i+1, -nums[i])
		if len(s) > 0 {
			for j := 0; j < len(s); j++ {
				ans = append(ans, append([]int{nums[i]}, s[j]...))
			}
		}
	}
	return ans
}

func twoSum(numbers []int, start, target int) [][]int {

	j := len(numbers) - 1
	ans := make([][]int, 0)

	for i := start; i < len(numbers); i++ {
		//两数之和去重
		if i > start && numbers[i] == numbers[i-1] {
			continue
		}

		for i < j && numbers[i]+numbers[j] > target {
			j--
		}
		if i < j && numbers[i]+numbers[j] == target {
			ans = append(ans, []int{numbers[i], numbers[j]})
		}
	}
	return ans
}
