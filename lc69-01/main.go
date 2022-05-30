package main

import "fmt"

/*
69. x 的平方根
https://leetcode.cn/problems/sqrtx/
*/

func mySqrt(x float64) float64 {

	left := 0.0
	right := x
	//终止于left=right
	for right-left > 1e-7 {
		// //+1 是上取整，避免只有两个元素时，产生死循环
		// mid := (left + right + 1) / 2
		//float 不需要取整，不会产生死循环，是一个精确的点
		mid := (left + right) / 2
		if mid*mid <= x {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func main() {

	fmt.Println(mySqrt(8.0))

}
