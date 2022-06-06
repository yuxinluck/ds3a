package main

import (
	"fmt"
)

func main() {

	l1 := []int{2, 5, 6, 7, 9, 5, 7}
	l2 := []int{1, 7, 4, 3, 4}

	ls := make([][]int, 0)
	ls = append(ls, l1)
	ls = append(ls, l2)
	k := 3
	// pos := 0

	// for i := 0; i < len(l); {
	// 	if i+k <= len(l) {
	// 		fmt.Println(l[i : i+k])
	// 	} else if i < len(l) {
	// 		fmt.Println(l[i:])
	// 	}
	// 	i = i + k
	// }

	ans := make([]int, 0)

	flag := true
	pos := 0

	for flag {
		flag = false
		for _, l := range ls {
			if pos >= len(l) {
				continue
			}
			if pos+k <= len(l) {
				ans = append(ans, l[pos:pos+k]...)
				flag = true // 只要有新数据，并且原数组还没用完，就继续
			} else if pos < len(l) {
				ans = append(ans, l[pos:]...)
			}
		}
		pos = pos + k
	}

	fmt.Println(ans)
}
