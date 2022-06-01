package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
某学校举行运动会,学生们按编号（1、2、3.....n)进行标识,
   现需要按照身高由低到高排列，
   对身高相同的人，按体重由轻到重排列，
   对于身高体重都相同的人，维持原有的编号顺序关系。
   请输出排列后的学生编号
   输入描述：
      两个序列，每个序列由N个正整数组成，(0<n<=100)。
      第一个序列中的数值代表身高，第二个序列中的数值代表体重，
   输出描述：
      排列结果，每个数据都是原始序列中的学生编号，编号从1开始，
   实例一：
      输入:
       4
       100 100 120 130
       40 30 60 50
      输出:
       2134
*/

func main() {

	h := []int{100, 100, 120, 130}
	w := []int{40, 30, 60, 50}

	n := len(h)

	res := make([][]int, 0)

	for i := 0; i < n; i++ {
		res = append(res, []int{h[i], w[i], i + 1})
	}
	fmt.Println(res)

	sort.Slice(res, func(i, j int) bool {
		for n := 0; n < 3; n++ {
			if res[i][n] < res[j][n] {
				return true
			}
		}
		return false
	})

	s := ""
	for i:=0;i<n;i++{
		
		s = s + strconv.Itoa(res[i][2])
	}
	fmt.Println(s)

}
