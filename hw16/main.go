package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
http://www.amoscloud.com/?p=2426
    单词接龙的规则是
    可用于接龙的单词 首字母必须要与前一个单词的尾字母相同
    当存在多个首字母相同的单词时，取长度最长的单词
    如果长度也相等，则取字典序最小的单词
    已经参与接龙的单词不能重复使用
    现给定一组全部由小写字母组成的单词数组
    并指定其中一个单词为起始单词
    进行单词接龙
    请输出最长的单词串
    单词串是单词拼接而成的中间没有空格

    输入描述
    输入第一行为一个非负整数
    表示起始单词在数组中的索引k
    0<=k<N
    输入的第二行为非负整数N
    接下来的N行分别表示单词数组中的单词

    输出描述，
    输出一个字符串表示最终拼接的单词串

    示例
    0
    6
    word
    dd
    da
    dc
    dword
    d

    输出
    worddwordda
    说明 先确定起始单词word 在接dword
    剩余dd da dc 则取da

   示例2
    4
    6
    word
    dd
    da
    dc
    dword
    d

    输出
    dwordda

    单词个数1<N<20
    单个单词的长度  1~30

*/

var src = `word
dd
da
dc
dword
d`

func main() {

	k := 0

	s := strings.Split(src, "\n")
	ans := ""
	ans += s[k]
	if k == 0 {
		s = s[1:]
	} else if k == len(s)-1 {
		s = s[0 : len(s)-1]
	} else {
		s = append(s[0:k], s[(k+1):]...)
	}

	tail := string(ans[len(ans)-1])
	//将s按照题意排序，先按首字母字典顺序 由小到大，再按照长度 由长到短，再按照字典顺序由小到大
	sort.Slice(s, func(i, j int) bool {
		if s[i][0] < s[j][0] {
			return true
		} else if s[i][0] == s[j][0] {
			if len(s[i]) > len(s[j]) {
				return true
			} else if len(s[i]) == len(s[j]) {
				if s[i] < s[j] {
					return true
				}
				return false
			}
			return false
		}
		return false
	})

	// fmt.Println(tail)
	// fmt.Println(s)
	idx := -1
	//当没有可接龙的字符串就终止
	for idx != len(s) {
		idx = sort.Search(len(s), func(i int) bool {
			return strings.HasPrefix(s[i], tail)
		})

		if idx != len(s) {
			//元素用掉就设置为空，不影响后续查询
			ans += s[idx]
			s[idx] = ""
			tail = string(ans[len(ans)-1])
		}
	}
	fmt.Println(ans)
}
