package main

import "fmt"

/*
http://www.amoscloud.com/?p=2391
题目描述：
一天一只顽猴想要从山脚爬到山顶
途中经过一个有n个台阶的阶梯，但是这个猴子有个习惯，每一次只跳1步或3步
试问？猴子通过这个阶梯有多少种不同的跳跃方式

输入描述
输入只有一个这个数n 0 < n < 50
此阶梯有多个台阶

输出描述
有多少种跳跃方式
*/

func main() {

	fmt.Println(solution(50))
}

//当前台阶的次数，为前 n-1的次数 + 前 n-3 的次数
func solution(n int) int {
	//一节台阶右一种方法
	step1 := 1
	//二节台阶右一种方法
	step2 := 1
	//三节台阶右一种方法
	step3 := 2
	//step
	step := 0
	if n == 1 || n == 2 {
		step = 1
	} else {
		step = 2
	}

	for i := 4; i <= n; i++ {
		step = step3 + step1
		step1 = step2
		step2 = step3
		step3 = step
	}
	return step
}
