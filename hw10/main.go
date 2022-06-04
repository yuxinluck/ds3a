package main

import "fmt"

/*
题目描述：
对一个数据a进行分类，
分类方法是，此数据a(4个字节大小)的4个字节相加对一个给定值b取模，
如果得到的结果小于一个给定的值c则数据a为有效类型，其类型为取模的值。
如果得到的结果大于或者等于c则数据a为无效类型。

比如一个数据a=0x01010101，b=3
按照分类方法计算：(0x01+0x01+0x01+0x01)%3=1
所以如果c等于2，则此a就是有效类型，其类型为1
如果c等于1，则此a是无效类型

又比如一个数据a=0x01010103，b=3
按分类方法计算：(0x01+0x01+0x01+0x03)%3=0
所以如果c=2则此a就是有效类型 其类型为0
如果c等于0 则此a是无效类型

输入12个数据，
第一个数据为c，第二个数据为b，
剩余10个数据为需要分类的数据

请找到有效类型中包含数据最多的类型，
并输出该类型含有多少个数据
*/

func main() {
	nums := []int{256, 257, 258, 259, 260, 261, 262, 263, 264, 265}
	b := 4
	c := 3

	solution(nums, b, c)
}

func solution(nums []int, b int, c int) {

	m := make(map[int]int)

	for _, v := range nums {
		reme := getByteSum(v) % b
		fmt.Println(reme, getByteSum(v))
		if reme < c {
			m[reme] += 1
		}
	}

	fmt.Println(m)

	max := 0

	for _, v := range m {
		if max < v {
			max = v
		}
	}

	fmt.Println(max)

}

//根据题意，算四子节和
func getByteSum(n int) int {
	sum := 0
	for i := 0; i < 4; i++ {
		sum += (n >> (i * 8)) % 16
	}
	return sum
}
