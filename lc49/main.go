package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func groupAnagrams(strs []string) [][]string {

	s := make(map[string][]string)

	for _,i := range strs {
		if _,ok := s[calcHash(i)]; ok {
			s[calcHash(i)] = append(s[calcHash(i)], i)
		} else {
			s[calcHash(i)] = []string{i}
		}
	}
    
	ans := make([][]string,0)

	for _,v := range s {
		ans = append(ans, v)
	}

	return ans
}

func calcHash(s string) string {
	l := strings.Split(s, "")
	// sort.Slice(l, func(i int, j int) bool {
	// 	return l[i] < l[j]
	// })
	sort.Strings(l)
	return strings.Join(l, "")
}

func main() {
    strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	s := groupAnagrams(strs)

	fmt.Println(s)
}
