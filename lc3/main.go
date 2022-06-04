package main

import "strings"

/*
https://leetcode.cn/problems/longest-substring-without-repeating-characters/
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func lengthOfLongestSubstring(s string) int {

    sl:= strings.Split(s,"")
    ans:=0
    m := make(map[string]int)
    n := len(sl)
    r:=-1
    
    for i:=0;i<n;i++{
        if i!=0 { //左边向右移动一个删除一个字符
            delete(m, sl[i-1])
        }

        for r+1 < n && m[sl[r+1]] == 0 {
            m[sl[r+1]] += 1
            r++ // 不断的移动右指针
        }
        //从i到r就是无重复字符串长度
        if r - i + 1 > ans {
            ans = r -i +1
        }
    }
    return ans

}