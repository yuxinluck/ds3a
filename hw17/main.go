package main

import (
	"fmt"
	"regexp"
	"sort"
)

/*
    给定一个字符串
    只包含大写字母
    求在包含同一字母的子串中
    长度第K长的子串
    相同字母只取最长的子串

    输入
     第一行 一个子串 1<len<=100
     只包含大写字母
     第二行为k的值

     输出
     输出连续出现次数第k多的字母的次数

     例子：
     输入
             AABAAA
             2
     输出
             1
       同一字母连续出现最多的A 3次
       第二多2次  但A出现连续3次

    输入
    AAAAHHHBBCDHHHH
    3
    输出
    2
//如果子串中只包含同一字母的子串数小于k

则输出-1

*/

func main() {
	src := "AAAAHHHBBCDHHHH"
	m := make(map[string]int, 0)
	k := 3

	for _, v := range src {
		m[string(v)] = 1
	}

	fmt.Println(m)

	for k, _ := range m {
		re, _ := regexp.Compile("[^" + k + "]+")
		tmp := re.Split(src, -1)
		fmt.Println(tmp)

		for _, v := range tmp {
			if m[k] < len(v) {
				m[k] = len(v)
			}
		}
	}
	fmt.Println(m)
	ans := make([]int, 0)
	for _, v := range m {
		ans = append(ans, v)
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})

	fmt.Println(ans)

	if len(ans) < k {
		fmt.Println(-1)
	} else {
		fmt.Println(ans[k-1])
	}
}
