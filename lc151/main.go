package main

/*
151. 颠倒字符串中的单词
https://leetcode.cn/problems/reverse-words-in-a-string/
*/

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	re, _ := regexp.Compile("[ ]+")
	s = strings.TrimSpace(s)
	s = re.ReplaceAllString(s, " ")

	l := strings.Split(s, " ")
	d := ""
	for i := len(l) - 1; i >= 1; i-- {
        d = d + l[i]+" "
	}

	d = d + l[0]

	return d
}

func main() {
	s := "   hello   word  kkk  "

	s = reverseWords(s)
	fmt.Println(s)
}
