package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
http://www.amoscloud.com/?p=3557

小组中每位都有一张卡片
  卡片是6位以内的正整数
  将卡片连起来可以组成多种数字
  计算组成的最大数字

  输入描述：
    ","分割的多个正整数字符串
    不需要考虑非数字异常情况
    小组种最多25个人

   输出描述：
     最大数字字符串

   示例一
     输入
      22,221
     输出
      22221

    示例二
      输入
        4589,101,41425,9999
      输出
        9999458941425101
*/

func main() {

	s := "4589,101,41425,9999"
	nums := make([][]int, 0)

	for _, n := range strings.Split(s, ",") {
		temp := make([]int, 0)
		for _, i := range strings.Split(n, "") {
			t, _ := strconv.Atoi(i)
			temp = append(temp, t)
		}
		nums = append(nums, temp)
	}

	fmt.Println(nums)

	sort.Slice(nums, func(i, j int) bool {
		l := Min(len(nums[i]), len(nums[j]))
		for n := 0; n < l; n++ {
			if nums[i][n] > nums[j][n] {
				return true
			}
		}
		return false
	})
	fmt.Println(nums)
	res := ""

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			res = res + strconv.Itoa(nums[i][j])
		}
	}
	fmt.Println(res)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
