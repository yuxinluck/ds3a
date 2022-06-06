package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	src := "1,2,5,-21,22,11,55,-101,42,8,7,32"

	nums := make([]int, 0)

	for _, v := range strings.Split(src, ",") {
		tmp, _ := strconv.Atoi(v)
		nums = append(nums, tmp)
	}
	fmt.Println(nums)

	/*
			请按照数组元素十进制最低位从小到大进行排序
		    十进制最低位相同的元素，相对位置保持不变
		    当数组元素为负值时，十进制最低为等同于去除符号位后对应十进制值最低位
	*/

	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i]%10)) < math.Abs(float64(nums[j]%10))
	})

	fmt.Println(nums)
}
