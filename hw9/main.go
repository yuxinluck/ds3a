package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
http://www.amoscloud.com/?p=2409
题目描述：
给定两个字符串，
从字符串2中找出字符串1中的所有字符，
去重并按照ASCII码值从小到大排列，

输入描述
字符范围满足ASCII编码要求，
输入字符串1长度不超过1024，
字符串2长度不超过100

输出描述
按照ASCII由小到大排序

示例一
输入

bach
bbaaccddfg
COPY
输出

abc
COPY

*/

func main() {
	s1 := "bach"
	s2 := "bbaaccddfg"

	m := make(map[string]bool)
	d := make([]string, 0)

	l := strings.Split(s1, "")

	for _, v := range l {
		m[v] = true
	}

	for k, _ := range m {
		if !strings.Contains(s2, k) {
			delete(m, k)
			continue
		}
		d = append(d, k)
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})

	fmt.Println(strings.Join(d, ""))
}
