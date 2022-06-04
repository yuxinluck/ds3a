package main

/*
https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/

0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//分析：约瑟环问题:动态规划
//状态方程:dp[n] = num;n表示含有多少个元素的圆环，，num为n个环游戏最后剩下的那个数
//状态转移方程:dp[n] = (dp[n-1] + m)%n,dp[n-1] = (dp[n-2]+m)%n-1
//......dp[2] = (dp[1]+m)%2,dp[1] = start（start为开始游戏的第一个数字）
//将问题划分为子问题,想要知道n个环游戏最后剩下的数字可变为先求
//n-1个环游戏最后剩下的数字,应为dp[n] = (dp[n-1]+m)%n
//递归解法
/*
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n
}
*/
//迭代解法
func lastRemaining(n int, m int) int {
	dp := 0 //题目从数字零开始
	for i := 2; i <= n; i++ {
		dp = (dp + m) % i
	}
	return dp
}
