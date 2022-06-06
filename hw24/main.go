package main

import (
	"fmt"
	"sort"
)

func main() {
	h := []int{100, 100, 120, 130}
	w := []int{40, 30, 60, 50}

	ans := make([]int, len(h))

	for i := 0; i < len(h); i++ {
		ans[i] = i + 1
	}

	sort.Slice(ans, func(i, j int) bool {
		if h[i] < h[j] {
			return true
		} else if h[i] == h[j] {
			if w[i] < w[j] {
				return true
			}
			return false
		}
		return false
	})

	fmt.Println(ans)
}

/*
   身高从低到高
   身高相同体重从轻到重
   体重相同维持原来顺序
   输入
   4
   100 100 120 130
   40 30 60 50
   输出：
   2 1 3 4
   输入
   3
   90 110 90
   45 60 45
   输出
   1 3 2
*/
