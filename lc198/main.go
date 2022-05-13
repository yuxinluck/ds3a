package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = -1e9
		dp[i][1] = -1e9
	}

	dp[0][0]=0

	// fmt.Println(dp)

	nums = append([]int{0}, nums...)

	for i := 1; i <= n; i++ {
		for j := 0; j < 2; j++ {
            dp[i][0] = Max(dp[i-1][0],dp[i-1][1])
			dp[i][1] = dp[i-1][0]+nums[i]
		}
	}
    return Max(dp[n][1],dp[n][0])
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
    nums := []int{2,7,9,3,1}
	fmt.Println(rob(nums))
}
