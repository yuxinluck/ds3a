package main

/*
https://leetcode.cn/problems/maximum-subarray/
53. 最大子数组和

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
*/

/*
题解

最大子序和


求 sum(l,r) = s[i] - s[l-1] 的最大值，固定 i, 即求 s[l-1]最小
*/

func maxSubArray(nums []int) int {

	//存储所有前缀和
	s := make([]int, len(nums)+1)
	s[0] = 0
	for i := 1; i <= len(nums); i++ {
		s[i] = s[i-1] + nums[i-1]
	}

	//存最小前缀和
	premin := s[0]
	ans := -1000000000
	for i := 1; i <= len(nums); i++ {
		ans = Max(ans, s[i]-premin)
		premin = Min(premin, s[i])
	}
	return ans
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
