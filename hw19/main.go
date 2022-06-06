package main

import (
	"fmt"
	"math"
	"strings"
)

/*
   删除字符串中出现次数最少的字符
   如果多个字符出现次数一样则都删除

   例子：
   输入
     abcdd
     字符串中只
    输出
     dd

   输入
     aabbccdd

   输出
     empty

     如果都被删除  则换为empty

*/

func main() {
    s := "abcdd"

	m := make(map[string]int)

	for _,v := range s {
		m[string(v)] +=1
	}
	// fmt.Println(m)

	min := math.MaxInt

	for _,v := range m {
		if min > v {
			min = v
		}
	}

	for k,v := range m {
		if v == min {
            s = strings.ReplaceAll(s,k,"")
		}
	}

	if s == ""{
		fmt.Println("empty")
	} else {
		fmt.Println(s)
	}
}
