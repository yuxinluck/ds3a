package main

import "fmt"

/*
状态：
dp数组含义：
dp[i][0] 表示第i天不持有股票的收益
dp[i][1] 表示第i天持有股票的收益


当第i天持有股票，即dp[i][1]的情况，可以由两个状态递推出来：
a. 第 i-1 天持有股票，第i天保持原状，所得收益就是i-1天持有股票收益 dp[i-1][1]
b. 第i天买入，第 i-1 天不持有股票，所得收益就是i-1天不持有股票收益减去第i天股票价格  dp[i-1][0]-prices[i]

当第i天不持有股票，即dp[i][0]的情况，可以由两个状态递推出来：
a. 第i天保持原状，则第 i-1 天持不有股票，，所得收益就是i-1天不持有股票收益 dp[i-1][0]
b. 第i天卖出，则第 i-1 天持有股票，所得收益就是i-1持有股票收益加上第i天股票价格  dp[i-1][1]+prices[i]

*/

func maxProfit(prices []int) int {

	//1. define dp
	dp := make([][]int,len(prices))
	for i := range prices {
		dp[i]=make([]int, 2)
	}
	//2. dp 初始化
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	
	//3. loop over all states
	for i:=1;i<len(prices);i++ {
		//4. copy decisions
		dp[i][0] = Max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1],dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
    // prices = append([]int{0},prices...)

	fmt.Println(maxProfit(prices))
}
