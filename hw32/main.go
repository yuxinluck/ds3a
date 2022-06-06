package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	M := 3
	N := 6
	src := "12 10 20 30 15 23"
	nums := make([]int, 0)
	for _, v := range strings.Split(src, " ") {
		tmp, _ := strconv.Atoi(v)
		nums = append(nums, tmp)
	}

	fmt.Println(nums)

	max := 0
	for i := 0; i < M; i++ {
		max += nums[i]
	}
	tmp := max

	for i := M; i < N; i++ {
		tmp = tmp + nums[i] - nums[i-M]
		if max < tmp {
			max = tmp
		}
	}

	fmt.Println(max)

}

/*
   有一个N个整数的数组
   和一个长度为M的窗口
   窗口从数组内的第一个数开始滑动
   直到窗口不能滑动为止
   每次滑动产生一个窗口  和窗口内所有数的和
   求窗口滑动产生的所有窗口和的最大值

   输入描述
   第一行输入一个正整数N
   表示整数个数  0<N<100000
   第二行输入N个整数
   整数取值范围   [-100,100]
   第三行输入正整数M
   M代表窗口的大小
   M<=100000 并<=N

   输出描述
   窗口滑动产生所有窗口和的最大值

   示例一
   输入
   6
   12 10 20 30 15 23
   3

   输出
   68
*/
