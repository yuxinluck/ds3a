package main

import (
	"fmt"
	"strings"
)

/*
http://www.amoscloud.com/?p=2404
题目描述：
输入一个英文文章片段，
翻转指定区域的单词顺序，
标点符号和普通字母一样处理，
例如输入字符串
I am a developer.
[0,3]
则输出
developer. a am I

输入描述
使用换行隔开3个参数
第一个参数为文章内容 即英文字符串
第二个参数为翻转起始单词下标，下标从0开始
第三个参数为结束单词下标

输出描述
翻转后英文文章片段每个单词之间以一个半角空格分割输出

示例一
输入

I am a developer.
0
3
COPY
输出

developer. a am I
*/

func main() {
    
	s := "I am a developer."
	solution(s,0,3)

}

func solution(Str string, l int, r int) {
	worlds := strings.Split(Str, " ")

	if r > len(worlds)-1 {
		r = len(worlds) - 1
	}

	if len(worlds) == 0 || l < 0 || r-l <= 0 {
		fmt.Println("EMPTY")
		return
	}

	for l < r {
		worlds[l], worlds[r] = worlds[r], worlds[l]
		l++
		r--
	}

	for i := 0; i < len(worlds); i++ {
		fmt.Print(worlds[i])
		if i != len(worlds)-1 {
			fmt.Print(" ")
		}
	}

}
