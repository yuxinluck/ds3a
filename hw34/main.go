package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	src := "2 1 5 4 3 3 9 2 7 4 6 2 15 4 2 4"
	nums := make([]int, 0)
	for _, tmp := range strings.Split(src, " ") {
		n, _ := strconv.Atoi(tmp)
		nums = append(nums, n)
	}

	fmt.Println(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)

	m := make(map[int]int)

	for _, v := range nums {
		m[v] += 1
	}

	fmt.Println(m)

	count := -1

	for _, v := range m {
		if count < v {
			count = v
		}
	}

	fmt.Println(count)

	for k, v := range m {
		if v != count {
			delete(m, k)
		}
	}

	newNums := make([]int, 0, len(m))
	for k := range m {
		newNums = append(newNums, k)
	}

	fmt.Println(newNums)

	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i] < newNums[j]
	})

	if len(newNums)%2 == 0 {
		fmt.Println((newNums[len(newNums)/2-1] + newNums[len(newNums)/2]) / 2)
	} else {
		fmt.Println(newNums[len(newNums)/2])
	}
}

/*
http://www.amoscloud.com/?p=2621
   1.众数是指一组数据中出现次数多的数
   众数可以是多个
   2.中位数是指把一组数据从小到大排列，最中间的那个数，
   如果这组数据的个数是奇数，那最中间那个就是中位数
   如果这组数据的个数为偶数，那就把中间的两个数之和除以2就是中位数
   3.查找整型数组中元素的众数并组成一个新的数组
   求新数组的中位数

   输入描述
   输入一个一维整型数组，数组大小取值范围   0<n<1000
   数组中每个元素取值范围，  0<e<1000

   输出描述
   输出众数组成的新数组的中位数

   示例一
   输入：
   10 11 21 19 21 17 21 16 21 18 16
   输出
   21

   示例二
   输入
   2 1 5 4 3 3 9 2 7 4 6 2 15 4 2 4
   输出
   3

    示例三
   输入
   5 1 5 3 5 2 5 5 7 6 7 3 7 11 7 55 7 9 98 9 17 9 15 9 9 1 39
   输出
   7
*/
