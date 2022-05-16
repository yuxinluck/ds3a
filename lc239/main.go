package main

import "fmt"

/*
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

/*
题解

维护双端队列单调性
*/

func maxSlidingWindow(nums []int, k int) []int {

	ans := make([]int, 0)

	//双端队列存储元素下标,递减队列，队首存最大值
	dq := make([]int, 0)

	for i := 0; i < len(nums); i++ {

		//删除出界元素
		if len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:]
		}
		//并维护队列单调性,插入新元素
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= nums[i] {
			dq = dq[0 : len(dq)-1]
		}
		dq = append(dq, i)
		//更新ans
		if i >= k-1 {
			ans = append(ans, nums[dq[0]])
		}
	}

	return ans
}

func main() {

	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	ans := maxSlidingWindow(nums, 3)
	fmt.Println(ans)
}
