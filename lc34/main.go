package main

/*
34. 在排序数组中查找元素的第一个和最后一个位置
https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
*/
func searchRange(nums []int, target int) []int {
	ans := make([]int, 0)
	/*
			[5,7,7,8,8,10]
			找第一个大于等于8，即 i=3
		    //后继， >= target
	*/

	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	ans = append(ans, right)

	/*
			[5,7,7,8,8,10]
			找第一个小于等于8，即 i=4
		    //后继， >= target
	*/

	left = -1
	right = len(nums) - 1
	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	ans = append(ans, left)

	if ans[0] > ans[1] {
		return []int{-1, -1}
	}
	return ans

}

func main() {

}
