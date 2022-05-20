package main

/*
50. Pow(x, n)
https://leetcode.cn/problems/powx-n/
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn ）。
*/

/*
分治思想
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	var temp float64 = myPow(x, n/2)
	var ans float64 = temp * temp
	if n%2 == 1 {
		ans = ans * x
	}
	return ans
}
