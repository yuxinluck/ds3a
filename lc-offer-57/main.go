package main

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findContinuousSequence(target int) [][]int {
	//前缀和
	s := make([]int, target+1)
	s[0] = 0
	for i := 1; i <= target-1; i++ {
		s[i] = s[i-1] + i
	}

	ans := make([][]int, 0)
	for i := 1; i <= target-2; i++ {
		for j := i + 1; j < target-1; j++ {
			if s[j]-s[i-1] == target {
				tmp := make([]int, 0)
				for m := i; m <= j; m++ {
					tmp = append(tmp, m)
				}
				ans = append(ans, tmp)
				break
			}
			if s[j]-s[i-1] > target {
				break
			}
		}
	}
	return ans
}
