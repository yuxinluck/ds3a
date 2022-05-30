package main

import "fmt"

/*
https://leetcode.cn/problems/binary-search/
704. 二分查找
*/

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {

	nums := []int{-1,0,3,5,9,12}

	fmt.Println(search(nums,9))

}
