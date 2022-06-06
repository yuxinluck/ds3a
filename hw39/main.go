package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	n := 2

	res := "1"

	if n == 0 {
		fmt.Println(res)
		return
	}

	for i := 1; i <= n; i++ {
		tmp := ""
		chars := strings.Split(res, "")
		prec := chars[0]
		count := 1
		for j := 1; j < len(chars); j++ {
			if chars[j] == prec {
				count++
			} else {
				tmp = tmp + strconv.Itoa(count) + prec
				count = 1
				prec = chars[j]
			}
		}
		res = tmp + strconv.Itoa(count) + prec
	}

	fmt.Println(res)

}

/*
   有一个数列A[n]
   从A[0]开始每一项都是一个数字
   数列中A[n+1]都是A[n]的描述
   其中A[0]=1
   规则如下
   A[0]:1
   A[1]:11 含义其中A[0]=1是1个1 即11
   表示A[0]从左到右连续出现了1次1
   A[2]:21 含义其中A[1]=11是2个1 即21
   表示A[1]从左到右连续出现了2次1
   A[3]:1211 含义其中A[2]从左到右是由一个2和一个1组成 即1211
   表示A[2]从左到右连续出现了一次2又连续出现了一次1
   A[4]:111221  含义A[3]=1211 从左到右是由一个1和一个2两个1 即111221
   表示A[3]从左到右连续出现了一次1又连续出现了一次2又连续出现了2次1

    输出第n项的结果
    0<= n <=59
    输入描述：
    数列第n项   0<= n <=59
    4
    输出描述
    数列内容
    111221
*/
