package main

import (
	"fmt"
	"sort"
)

func main() {
	src := []int{2, 7, 3, 0}

	m := make(map[int]int)
	for _, v := range src {
		m[v] = 1
	}

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	ans := make([]int, 0)
	for a := len(keys) - 1; a > 0; a-- {
		for b := 0; b < a; b++ {
			for c := 0; c < a; c++ {
				if keys[a] == keys[b]+2*keys[c] && b != c {
					ans = append(ans, keys[a], keys[b], keys[c])
				}
			}
		}
	}
    json.u
	if len(ans) == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(ans)
	}
}
