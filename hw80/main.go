package main

import (
	"fmt"
	"strconv"
)

/*
     给定参数n,从1到n会有n个整数:1,2,3,...,n,
     这n个数字共有n!种排列.
   按大小顺序升序列出所有排列的情况,并一一标记,
   当n=3时,所有排列如下:
   "123" "132" "213" "231" "312" "321"
   给定n和k,返回第k个排列.

   输入描述:
     输入两行，第一行为n，第二行为k，
     给定n的范围是[1,9],给定k的范围是[1,n!]。
   输出描述：
     输出排在第k位置的数字。

   实例1：
     输入:
       3
       3
     输出：
       213
     说明
       3的排列有123,132,213...,那么第三位置就是213

   实例2：
     输入
       2
       2
     输出：
       21
     说明
       2的排列有12,21，那么第二位置的为21
*/

/*
题解
递归
从小到大全排列
*/

var n int
var k int

func Permutation(nums []int) int {

	ans := make([][]int, 0)

	temp := make([]int, 0)

	selected := make([]bool, n)

	var choice func(nums []int, temp []int, selected []bool)

	choice = func(nums []int, temp []int, selected []bool) {
		if len(temp) >= n {
			ans = append(ans, temp)
			if len(ans) >= k {
				return
			}
		}

		for i := 0; i < n; i++ {
			if !selected[i] {
				temp = append(temp, nums[i])
				selected[i] = true
				choice(nums, temp, selected)
				temp = temp[0 : len(temp)-1]
				selected[i] = false
			}
		}
	}
	choice(nums, temp, selected)

	fmt.Println(ans)
	// fmt.Println(ans[k-1])
	s := ""
	for i := 0; i < n; i++ {
		s = s + strconv.Itoa(ans[k-1][i])
	}
	fmt.Println(s)

	res, _ := strconv.Atoi(s)
	return res
}

func main() {

	fmt.Scanln(&n)
	fmt.Scanln(&k)

	// fmt.Printf("%T , %d ", n, n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// fmt.Println(nums)

	fmt.Println(Permutation(nums))
}
