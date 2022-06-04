package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
题目描述：
小明今年升学到了小学一年级，
来到新班级后，发现其他小朋友身高参差不齐，
然后就想基于各小朋友和自己的身高差，对他们进行排序，
请帮他实现排序

输入描述
第一行为正整数 H和N
0 < H < 200 为小明的身高
0 < N < 50 为新班级其他小朋友个数
第二行为N个正整数
H1 ~ Hn分别是其他小朋友的身高
取值范围0 < Hi < 200
且N个正整数各不相同

输出描述
输出排序结果，各正整数以空格分割
和小明身高差绝对值最小的小朋友排在前面
和小明身高差绝对值最大的小朋友排在后面
如果两个小朋友和小明身高差一样
则个子较小的小朋友排在前面
*/

func main() {
	s := "95 96 97 98 99 101 102 103 104 105"
	highs := make([]int, 0)
	for _, v := range strings.Split(s, " ") {
		n, _ := strconv.Atoi(v)
		highs = append(highs, n)
	}

	solution(100, highs)
}

func solution(h int, highs []int) {

	//根据题意，排序条件
	sort.Slice(highs, func(i, j int) bool {
		if IntAbs(highs[i], h) < IntAbs(highs[j], h) {
			return true
		} else if IntAbs(highs[i], h) == IntAbs(highs[j], h) {
			return highs[i] < highs[j]
		}
		return false
	})

	fmt.Println(highs)
}

func IntAbs(x int, y int) int {
	if x-y >= 0 {
		return x - y
	}
	return y - x
}
