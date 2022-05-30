package main

import "fmt"

/*
https://leetcode.cn/problems/minimum-number-of-days-to-make-m-bouquets/
1482. 制作 m 束花所需的最少天数
*/

func minDays(bloomDay []int, m int, k int) int {
	n := m * k
	if len(bloomDay) < n {
		return -1
	}

	l := 1
	r := 1

	for i := 0; i < len(bloomDay); i++ {
		r = Max(r, bloomDay[i])
	}

	for l < r {
		mid := (l + r) / 2
		cnt := 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= mid {
				cnt++
			}
		}

		if cnt < n {
			l = mid+1
		} else {
			if isVaild(bloomDay,mid, m, k) {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}
	return r
}

func isVaild(nums []int, mid int, m int, k int) bool {
	cnt := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <=mid {
			temp++
			if temp == k {
				cnt++
				temp = 0
			}
		} else {
			temp = 0
		}
	}
	return cnt >= m
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {

	bloomDay := []int{1, 10, 3, 10, 2}
	r := minDays(bloomDay, 3, 1)
	fmt.Println(r)
}
