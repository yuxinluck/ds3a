package main

import "fmt"

/*
题目描述：
给航天器一侧加装长方形和正方形的太阳能板(图中的斜线区域);
需要先安装两个支柱(图中的黑色竖条);
再在支柱的中间部分固定太阳能板;
但航天器不同位置的支柱长度不同;
太阳能板的安装面积受限于最短一侧的那支支柱的长度;

现提供一组整型数组的支柱高度数据;
假设每个支柱间的距离相等为一个单位长度;
计算如何选择两根支柱可以使太阳能板的面积最大;

输入描述
10,9,8,7,6,5,4,3,2,1
注释，支柱至少有两根，最多10000根，能支持的高度范围1~10^9的整数

柱子的高度是无序的
例子中的递减是巧合

输出描述
可以支持的最大太阳板面积:(10m高支柱和5m高支柱之间)
25

示例一
输入

10,9,8,7,6,5,4,3,2,1
COPY
输出

25
COPY
备注
10米高支柱和5米高支柱之间宽度为5，高度取小的支柱高度也是5，面积为25
任取其他两根支柱所能获得的面积都小于25 所以最大面积为25

/
最大矩形
/

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
			ans = Max(ans, s[len(s)-1].Height*(accumulateWidth-1))
			s = s[:len(s)-1]
		}
		s = append(s, Rect{h, accumulateWidth + 1})
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
	heights := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	area := largestRectangleArea(heights)

	fmt.Println(area)

}
