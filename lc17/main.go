package main

import (
	"strings"
)

/*
17. 电话号码的字母组合
https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
*/

func letterCombinations(digits string) []string {

	if digits == "" {
		return nil
	}
	digs := strings.Split(digits, "")
	alphapet := make(map[string]string)
	alphapet["2"] = "abc"
	alphapet["3"] = "def"
	alphapet["4"] = "ghi"
	alphapet["5"] = "jkl"
	alphapet["6"] = "mno"
	alphapet["7"] = "pqrs"
	alphapet["8"] = "tuv"
	alphapet["9"] = "wxyz"

	substr := make([]string, 0)
	ans := make([]string, 0)

	var dfs func(index int)

	dfs = func(index int) {

		if index >= len(digs) {
			ans = append(ans, strings.Join(substr, ""))
			return
		}
		for _, s := range alphapet[digs[index]] {
			substr = append(substr, string(s))
			dfs(index + 1)
			substr = substr[0 : len(substr)-1]
		}
	}

	dfs(0)

	return ans
}

func main() {
	letterCombinations("123")
}
