package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
题目描述：
一辆运送快递的货车，
运送的快递放在大小不等的长方体快递盒中，
为了能够装载更多的快递同时不能让货车超载，
需要计算最多能装多少个快递。
注：快递的体积不受限制。
快递数最多1000个，货车载重最大50000。

输入描述
第一行输入每个快递的重量
用英文逗号隔开
如 5,10,2,11
第二行输入货车的载重量
如 20

输出描述
输出最多能装多少个快递
如 3

示例一
输入

5,10,2,11
20
COPY
输出

3
*/

func main() {

	s1 := "5,10,2,11"
	s2 := "20"

	car, _ := strconv.Atoi(s2)
	ws := make([]int, 0)

	for _, v := range strings.Split(s1, ",") {
		t, _ := strconv.Atoi(v)
		ws = append(ws, t)
	}
	sort.Slice(ws, func(i, j int) bool { return ws[i] < ws[j] })
	sum := 0
	account := 0

	for i := 0;i<len(ws);i++{
		sum += ws[i] 
		if sum > car {
			fmt.Println(account)
			return
		} 
		account++
	}

}
