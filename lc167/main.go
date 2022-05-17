package main

/*
167. 两数之和 II - 输入有序数组
https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
*/

/*
双指针
固定左边，找右边
*/

func twoSum(numbers []int, target int) []int {

	j := len(numbers) - 1
	for i := 0; i < len(numbers); i++ {
		for i < j && numbers[i]+numbers[j] > target {
			j--
		}
		if i < j && numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return nil

}
