package main

/*
154. 寻找旋转排序数组中的最小值
https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/
*/

/*
题解
[4,5,6,7,0,1,2]

二分条件：
将每个元素与最后一个元素比较，<= 为true,否则为false
[0,0,0,0,1,1,1]
*/

func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return nums[left]
}

func main() {

}
