package main

import (
	"fmt"
	"reflect"
)

/*
30.串联所有单词的子串
给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/substring-with-concatenation-of-all-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
题解：
将s分成n个长度为 len(words)*len(word[0]) 的子段；
比较子段，words map 是否相等

*/

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)

	//wordMp
	wordsMap := make(map[string]int)
	//单个word长度
	k := len(words[0])
	//子段长度
	tot := len(words)*k
	for _, v := range words {
		wordsMap[v]++
	}

	// fmt.Println(wordsMap, tot)
	//分解s
	for i := 0; i+tot <= len(s); i++ {
        // fmt.Println(s[i:i+tot])
		if valid(s[i:i+tot],k,wordsMap){
			ans=append(ans, i)
		}
	}

	return ans
}

func valid(sub string, k int, wordsMap map[string]int ) bool {
	subMap := make(map[string]int)

    for i:=0;i+k<=len(sub);i+=k {
        subMap[sub[i:i+k]]++
	}
    
	return reflect.DeepEqual(subMap,wordsMap)
}

func main() {

	words := []string{"foo", "bar"}
	s := findSubstring("barfoothefoobarman", words)

	fmt.Println(s)
}
