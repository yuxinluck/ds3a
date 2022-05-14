package main

/*
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			if c == ')' {
				if stack[len(stack)-1] != '(' {
					return false
				}
			} else if c == ']' {
				if stack[len(stack)-1] != '[' {
					return false
				}
			} else {
				if stack[len(stack)-1] != '{' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
