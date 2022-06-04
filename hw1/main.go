package main

import (
	"fmt"
	"math"
)

/*
题目描述：
如果三个正整数A、B、C ,A²+B²=C²则为勾股数
如果ABC之间两两互质，即A与B，A与C，B与C均互质没有公约数，
则称其为勾股数元组。
请求出给定n~m范围内所有的勾股数元组

输入描述
起始范围
1 < n < 10000
n < m < 10000

输出描述
ABC保证A<B<C
输出格式A B C
多组勾股数元组，按照A B C升序的排序方式输出。
若给定范围内，找不到勾股数元组时，输出Na。
*/

func solution(n int, m int) {
	count := 0

	for a := n; a < m-1; a++ {
		for b := a + 1; b < m; b++ {
			for c := b + 1; c < m+1; c++ {
				if isPreme(a,b)&&isPreme(a,c)&&isPreme(b,c)&& a*a+b*b==c*c {
					count++
					fmt.Println(a, b, c)
				}
			}
		}
	}

	if count == 0 {
		fmt.Println("Na")
	}
}

//判断是否护质
func isPreme(x, y int) bool {
	min := Min(x, y)
	s := math.Sqrt(float64(min))
	for i := 2; float64(i) < s; i++ {
		if x%i == 0 && y%i == 0 {
			return false
		}
	}
	return true
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main(){
	solution(5,10)
}