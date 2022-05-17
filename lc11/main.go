package main

/*
11. 盛最多水的容器
https://leetcode.cn/problems/container-with-most-water/
*/

/*
题解
双指针，宽度一定，哪边小移动拿边

暴力解法，简单明了好想通，单一般是最差解法；
努力寻找最优解，减少循环层数，减少计算量
*/

func maxArea(height []int) int {
	ans := 0

	i := 0
	j := len(height) - 1

	for i < j {
		ans = Max(ans, Min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
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
