package main

/*
1，声明一个int型变量depth和整型数组变量ans
2，遍历字符串seq，当字符为(时深度递增，同时将深度值append进ans；否则字符就是)，将字符为)的深度值append进ans，同时深度值-1
3，返回最终结果ans

作者：zoctopus_zhang
链接：https://leetcode.cn/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/solution/guan-fang-ti-jie-fang-fa-yi-golangban-ben-by-zocto/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func maxDepthAfterSplit(seq string) []int {
	depth := 0
	ans := make([]int, 0)

	for _, v := range seq {
		if v == '(' {
			depth++
			ans = append(ans, depth%2)
		} else {
			ans = append(ans, depth%2)
			depth--
		}
	}
	return ans
}
