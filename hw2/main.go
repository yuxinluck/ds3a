package main

import (
	"fmt"
	"sort"
)

/*
http://www.amoscloud.com/?p=2387

给定两个整数数组，arr1、arr2，数组元素按升序排列；
假设从arr1、arr2中分别取出一个元素，可构成一对元素；
现在需要取出k对元素，并对取出的所有元素求和，计算和的最小值；
注意：两对元素对应arr1、arr2的下标是相同的，视为同一对元素。

输入描述
输入两行数组arr1、arr2
每行首个数字为数组大小size， 0 < size <= 100
arr1，arr2中的每个元素e， 0< e <1000
接下来一行，正整数k 0 < k <= arr1.size * arr2.size

输出描述
满足要求的最小值
*/

func main() {

	l1 := []int{1, 1, 2}
	l2 := []int{1, 2, 3}
	k := 2

	solution(l1, l2, k)

}

func solution(l1 []int, l2 []int, k int) {

	//求所有对数的和
	sums := make([]int, 0)

	for _, x := range l1 {
		for _, y := range l2 {
			sums = append(sums, x+y)
		}
	}

	//对所有和进行排序
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] < sums[j]
	})

	//输出sums的前k对和

	ans := 0
	for i := 0; i < k; i++ {
		ans += sums[i]
	}
	fmt.Println(ans)
}
