package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
   在通信系统中有一个常见的问题是对用户进行不同策略的调度
   会得到不同系统消耗的性能
   假设由N个待串行用户，每个用户可以使用A/B/C三种不同的调度策略
   不同的策略会消耗不同的系统资源
   请你根据如下规则进行用户调度
   并返回总的消耗资源数
   规则是：相邻的用户不能使用相同的调度策略
   例如：
   第一个用户使用A策略 则第二个用户只能使用B和C策略
   对单的用户而言，不同的调度策略对系统资源的消耗可以规划后抽象为数值
   例如
   某用户分别使用ABC策略的系统消耗，分别为15 8 17
   每个用户依次选择当前所能选择的对系统资源消耗最少的策略,局部最优
   如果有多个满足要求的策略，选最后一个

   输入描述：
       第一行表示用户个数N
       接下来表示每一行表示一个用户分别使用三个策略的资源消耗
       resA resB resC

   输出描述：
       最优策略组合下的总的系统消耗资源数

    示例一：
     输入：
         3
         15 8 17
         12 20 9
         11 7 5
     输出：
         24
      说明:
       1号用户使用B策略
       2号用户使用C策略
       3号用户使用B策略
      系统资源消耗8+9+7

*/

var s = `15 8 17
12 20 9
11 7 5`

func main() {

	src := strings.Split(s, "\n")
	fmt.Println(len(src))
	res := make([](map[int]int), 0)

	for _, v := range src {
		l := strings.Split(v, " ")
		tmp := make(map[int]int)
		for i := 0; i < 3; i++ {
			t, _ := strconv.Atoi(l[i])
			tmp[i] = t
		}
		res = append(res, tmp)
	}

	fmt.Println(res)

	ans := make([]int, 0)

	//第一行选择最小的
	ans1 := math.MaxInt
	abcType := 0
	for k, v := range res[0] {
		if ans1 > v {
			ans1, abcType = v, k
		}
	}
	fmt.Println(abcType)
	ans = append(ans, ans1)

	//根据题意条件，进行后续行的选择
	for i := 1; i < len(res); i++ {
		if abcType == 0 {
			if res[i][2] <= res[i][1] {
				abcType = 2
				ans = append(ans, res[i][2])
			} else {
				abcType = 1
				ans = append(ans, res[i][1])
			}
			continue
		}

		if abcType == 1 {
			if res[i][2] <= res[i][0] {
				abcType = 2
				ans = append(ans, res[i][2])
			} else {
				abcType = 0
				ans = append(ans, res[i][0])
			}
			continue
		}

		if abcType == 2 {
			if res[i][1] <= res[i][0] {
				abcType = 1
				ans = append(ans, res[i][1])
			} else {
				abcType = 0
				ans = append(ans, res[i][0])
			}
			continue
		}
	}

	fmt.Println(ans)

	sum := 0
	for _, v := range ans {
		sum += v
	}
	fmt.Println(sum)
}
