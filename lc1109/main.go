package main

/*
这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/corporate-flight-bookings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
题解：
差分思想

使用差分数组与原数组的关系，差分数组的前缀和就是原数组，求原数组
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	/*
		为了明确思路，写的因果
		//初始数组,默认全部为零
		initA := make([]int,n)
		//差分数组
		delta := make([]int,n+1)
		//差分数组关系
		delta[0]=initA[0]
		for i:=1;i<len(initA);i++{
			delta[i]=initA[i]-initA[i-1]
		}

		fmt.Println(initA,delta)
	*/
	//初始数组值为零，长度为n, 所以初始差分数组值为零，长度为n
	delta := make([]int, n+1)

	for _, v := range bookings {
		l := v[0] - 1
		r := v[1] - 1
		n := v[2]

		//原数组[l-r] +n ,差分数组 [l]+n, [r+1] -n
		delta[l] += n
		delta[r+1] -= n
	}

	//差分数组前缀和即位原数组当前状态
	sum := make([]int, n+2)
	sum[0] = 0
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + delta[i-1]
	}
	return sum[1 : n+1]

}

func main() {
	corpFlightBookings([][]int{[]int{}}, 5)
}
