package main

/*
69. x 的平方根
https://leetcode.cn/problems/sqrtx/
*/

func mySqrt(x int) int {

	left := 1
	right := x
	//终止于left=right
	for left < right {
		//+1 是上取整，避免只有两个元素时，产生死循环
		mid := (left + right + 1) / 2
		if mid*mid <= x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {

}
