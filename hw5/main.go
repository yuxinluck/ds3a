package main

import "fmt"

/*
http://www.amoscloud.com/?p=2393
题目描述：
为了充分发挥Gpu算力，
需要尽可能多的将任务交给GPU执行，
现在有一个任务数组，
数组元素表示在这1s内新增的任务个数，
且每秒都有新增任务，
假设GPU最多一次执行n个任务，
一次执行耗时1s，
在保证Gpu不空闲的情况下，最少需要多长时间执行完成。

输入描述
第一个参数为gpu最多执行的任务个数
取值范围1~10000
第二个参数为任务数组的长度
取值范围1~10000
第三个参数为任务数组
数字范围1~10000

输出描述
执行完所有任务需要多少秒
*/

func main() {

	n := 3
	jobs := []int{1,2,3,4,5}

	fmt.Println(solution(n,jobs))

}

func solution(n int, jobs []int) int {
    t := 0
	remaining := 0

	for _,j := range jobs {
		if j+remaining > n {
			remaining = j+remaining - n
		}else {
			remaining = 0
		}
		t ++
	}
    
	t += remaining/n

	if remaining%n > 0 {
		t++
	}  
	return t
}
