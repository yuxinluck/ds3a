package main

import (
	"fmt"
	"strings"
)

/*
https://www.nowcoder.com/exam/oj/ta?difficulty=2&page=1&pageSize=100&search=&tpId=37&type=37

HJ11
输入一个整数，将这个整数以字符串的形式逆序输出
程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001


数据范围： 0 \le n \le 2^{30}-1 \0≤n≤2
30
 −1
*/

func main() {
    var num string
	fmt.Scan(&num)

	l:= strings.Split(num, "")

	d:=""

	for i:=len(l)-1;i>=0;i--{
		d += l[i]
	}

	fmt.Println(d)
}
