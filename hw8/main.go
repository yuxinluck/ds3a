package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
http://www.amoscloud.com/?p=2406
题目描述：
双11众多商品进行打折销售，小明想购买一些自己心仪的商品，
但由于受购买资金限制，所以他决定从众多心意商品中购买3件，
而且想尽可能的花完资金，
现在请你设计一个程序帮助小明计算尽可能花费的最大资金额。

输入描述
第一行为整型数组M，数组长度小于100，数组元素记录单个商品的价格；
单个商品价格小于1000；
第二行输入为购买资金的额度R；
R < 100000。

输出描述
输出为满足上述条件的最大花费额度
如果不存在满足上述条件的商品请返回-1

示例一
输入

23,26,36,27
78
COPY
输出

76
COPY
*/

func main() {
	m := "23,26,28,30"

	solution(m, 80)

}

func solution(m string, r int) {
	prices := make([]int, 0)
	for _, s := range strings.Split(m, ",") {
		t, _ := strconv.Atoi(s)
		prices = append(prices, t)
	}

	if len(prices) < 3 || (len(prices) == 3 && GetSum(prices) > r) {
		fmt.Println(-1)
		return
	}

	//三件商品的总价
	sums := make([]int, 0)

	//三件商品组合

	var pailie func(tmp []int, start int)

	//递归，排列组合，获取不重复的商品组合的总价
	pailie = func(tmp []int, start int) {
		if len(tmp) == 3 {
			t := GetSum(tmp)
			sums = append(sums, t)
			return
		}
		for i := start; i < len(prices); i++ {
			tmp = append(tmp, prices[i])
			pailie(tmp, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
    pailie([]int{},0)

	fmt.Println(sums)
	//对sums排序
	sort.Slice(sums,func(i, j int) bool {
		return sums[i] < sums[j]
	})

	d := sort.Search(len(sums),func(i int) bool {
		return sums[i] >= r
	})
    
	if d == len(sums){
		fmt.Println(sums[len(sums)-1])
		return
	} 

	if sums[d] == r {
	    fmt.Println(sums[d])
		return
	} else {
		if d == 0 {
			fmt.Println(-1)
		    return
		}
		fmt.Println(sums[d-1])
		return
	}
}

//切片元素求和
func GetSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
