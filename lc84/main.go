package main

import "fmt"

/*
84. 柱状图中最大的矩形
https://leetcode.cn/problems/largest-rectangle-in-histogram/
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

*/

func largestRectangleArea(heights []int) int {
	//保存最大面积值
	ans := 0
	//记录矩形的栈
	s := make([]Rect, 0)
	//利用最后的零，保证最后栈被弹空
	heights = append(heights, 0)

	for _, h := range heights {
		//每列累计宽度
		accumulateWidth := 0
		//栈顶之前的 高度 >= 当前高度时，单调性被破坏，确定了栈顶高度的扩展范围，需要删除栈顶
		for len(s) != 0 && s[len(s)-1].Height >= h {
			accumulateWidth += s[len(s)-1].Width
			ans = Max(ans,s[len(s)-1].Height*accumulateWidth)
			s = s[:len(s)-1]
		}
        s = append(s, Rect{h,accumulateWidth+1})
	}

	return ans
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Rect struct {
	Height int
	Width  int
}

func main() {
    heights := []int{2,1,5,5,6,2,3}

	area := largestRectangleArea(heights) 

	fmt.Println(area)

}
