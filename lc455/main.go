package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {

	sort.Slice(g, func(i, j int) bool { return g[i] < g[j] })
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	j := 0
	ans := 0

	for _, child := range g {
		for j < len(s) && s[j] < child {
			j++
		}
		if j < len(s) {
			ans++
			j++
		}
	}

	return ans
}

func main() {

	g := []int{1, 2, 3}
	s := []int{1, 1}
    fmt.Println(findContentChildren(g,s))
}
