package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	n := 9
	ans := make([][]string, 0)
	ans = append(ans, []string{strconv.Itoa(n)})

	for i := 1; i < n; i++ {
		tmp := make([]string, 0)
		sum := 0
		for j := i; j < n; j++ {
			sum += j
			tmp = append(tmp, strconv.Itoa(j))
			if sum == n {
				ans = append(ans, tmp)
				break
			} else if sum > n {
				break
			}
		}
	}
	fmt.Println(ans)
	sort.Slice(ans, func(i, j int) bool {
		return len(ans[i]) < len(ans[j])
	})
	fmt.Println(ans)

	for _, v := range ans {
		fmt.Printf("%d=%s\n", n, strings.Join(v, "+"))
	}
	fmt.Println("Result:", len(ans))
}

/*
   一个整数可以由连续的自然数之和来表示
   给定一个整数
   计算该整数有几种连续自然数之和的表达式
   并打印出每一种表达式

   输入描述
   一个目标整数t  1<= t <=1000

   输出描述
   1.该整数的所有表达式和表达式的个数
   如果有多种表达式，自然数个数最少的表达式优先输出
   2.每个表达式中按自然数递增输出

   具体的格式参见样例
   在每个测试数据结束时，输出一行"Result:X"
   其中X是最终的表达式个数

   输入
   9

   输出
   9=9
   9=4+5
   9=2+3+4
   Result:3

   说明 整数9有三种表达方法：

   示例二
   输入
   10
   输出
   10=10
   10=1+2+3+4
   Result:2

*/
