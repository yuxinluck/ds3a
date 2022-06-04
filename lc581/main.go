package main

/*
581. 最短无序连续子数组

给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/shortest-unsorted-continuous-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findUnsortedSubarray(nums []int) int {
	ans := 0
	l := -1
	r := len(nums)

	flag := true
	//第一次遍历，找左边界
	for i := 0; i < len(nums); i++ {
		for l >= 0 && nums[i] < nums[l] {
			flag = false
			l--
		}
		if flag {
			l++
		}
	}
	//此时l在左侧有序序列的最后一位
	l++

	flag = true
	//第二次遍历，找右边界
	for i := len(nums) - 1; i >= 0; i-- {
		for r <= len(nums)-1 && nums[i] > nums[r] {
			flag = false
			r++
		}
		if flag {
			r--
		}
	}
	//此时r 在右侧有序序列的第一位
	r--

	if r > l {
		return r - l + 1
	}

	return ans
}
