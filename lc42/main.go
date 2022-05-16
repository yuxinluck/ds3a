package main

import "fmt"

/*

42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func trap(height []int) int {

	ans := 0
	s := make([]Rect, 0)

	for _, h := range height {
		//累计宽度
		accumulateWidth := 0
		//破坏单调性
		for len(s) > 0 && s[len(s)-1].Height <= h {
			bottom := s[len(s)-1].Height
			accumulateWidth += s[len(s)-1].Width
			s = s[:len(s)-1]
			if len(s) == 0 {
				continue
			}
			up := Min(s[len(s)-1].Height, h)
			ans += (up - bottom) * accumulateWidth
		}
		s = append(s, Rect{h, accumulateWidth + 1})
	}
	return ans
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type Rect struct {
	Height int
	Width  int
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	area := trap(height)

	fmt.Println(area)
}
