package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	src := "95 88 83 64 100"
	N := 2

	nums := make([]int, 0)

	m := make(map[string]bool)

	for _, v := range strings.Split(src, " ") {
		m[v] = true
	}

	if len(m) < 2*N {
		fmt.Println(-1)
		return
	}

	for k, _ := range m {
		tmp, _ := strconv.Atoi(k)
		nums = append(nums, tmp)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sum := 0

	for i := 0; i < N; i++ {
		sum += nums[i]
	}

	for j := len(nums) - N; j < len(nums); j++ {
		sum += nums[j]
	}

	fmt.Println(sum)
}

/*
   给定一个数组
   编写一个函数
   来计算他的最大N个数和最小N个数的和
   需要对数组进行去重

   说明
   第一行输入M
   M表示数组大小
   第二行输入M个数
   表示数组内容
   第三行输入N表示需要计算的最大最小N的个数

   输出描述
   输出最大N个数和最小N个数的和

   例一：
       输入
       5
       95 88 83 64 100
       2

       输出
       342

       说明
       最大2个数[100 95] 最小2个数[83 64]
       输出342

    例二
       输入
       5
       3 2 3 4 2
       2

       输出
        -1
        说明
        最大两个数是[4 3]最小2个数是[3 2]
        有重叠输出为-1

*/
