package main

import "fmt"

/*
https://leetcode.cn/problems/find-peak-element/
162. 寻找峰值
*/

/*
三分查找
*/

func findPeakElement(nums []int) int {

	l := 0
	r := len(nums) - 1

	for l < r {
		lmid := (l + r) / 2
		rmid := lmid + 1

		if nums[lmid] <= nums[rmid] {
			l = lmid + 1
		} else {
			r = rmid - 1
		}
	}
	return r
}

func main() {

	nums := []int{1,2,1,3,5,6,4}
    fmt.Println(findPeakElement(nums))
}
