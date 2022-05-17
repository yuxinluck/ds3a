package main

/*
1248. 统计「优美子数组」
给你一个整数数组 nums 和一个整数 k。如果某个连续子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。

请返回这个数组中 「优美子数组」 的数目。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/count-number-of-nice-subarrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func numberOfSubarrays(nums []int, k int) int {

	//前缀和数组
	s := make([]int, len(nums)+1)
	s[0] = 0
	//递推前缀和
	for i := 1; i <= len(nums); i++ {
		s[i] = s[i-1] + nums[i-1]%2 //奇偶转化为10
	}

	//问题转化为 sum(l-r)=s[r]-s[l-1]=k
	//固定 r ，k, 求 s[l-1]的个数
	ans := 0

	//根据题复杂度，开数组，或map记录s[l-1]
	//此题最大和为n，复杂度不高，开数组就够,索引为前缀和的值
	c := make([]int, len(nums)+1)
	c[s[0]]++

	for i := 1; i <= len(nums); i++ {
		if s[i]-k >= 0 {
			ans += c[s[i]-k]
		}
		c[s[i]]++
	}
	return ans

}
