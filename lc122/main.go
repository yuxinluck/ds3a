package main

import "fmt"

/*
这是一个预言家炒股的状态--知道未来每天的价格

当前持有股票，卖不卖？ 往后看一天，明天涨不卖，明天跌就卖
当前没有股票，买不买？往后看一天，明天涨就买，明天跌不买

最后就是最完美的结果 -- 获得所有 prices[i]-prices[i-1] > 0 区间的收益
*/

func maxProfit(prices []int) int {

	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += Max(prices[i]-prices[i-1], 0)
	}
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}
