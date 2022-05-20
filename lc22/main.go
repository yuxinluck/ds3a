package main

/*
22. 括号生成
https://leetcode.cn/problems/generate-parentheses/
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/

/*
题解

分治算法
一共n对括号
生成的组合 S，将S分为两部分 (A) 和 B , 第一段 (A) 外层必定是一对括号

(A)B

若 (A) 有k 对括号， 则 A ： k-1， B ：n-k

n = 3

k = 1 ; (A)  A=0 ; B=n-1=2  ()() (())
结果： ()()()   ()(())

k = 2 ; (A) A=1 () ; B=n-2=1 ()
结果： (())()

k = 3 ; (A) A=2 ()()  (()) ; B=n-3=0
结果：(()())  ((()))
*/

//分治，记录已经处理过的情况
var record map[int][]string = make(map[int][]string)

func generateParenthesis(n int) []string {

	ans := make([]string, 0)
	if n == 0 {
		return []string{""}
	}
	//如果已经有计算过的结果，直接返回结果
	if v, ok := record[n]; ok {
		return v
	}
	//加法原理
	for k := 1; k <= n; k++ {
		A := generateParenthesis(k - 1)
		B := generateParenthesis(n - k)

		//乘法原理
		for _, a := range A {
			for _, b := range B {
				ans = append(ans, "("+a+")"+b)
			}
		}
	}
	record[n] = ans
	return ans
}
