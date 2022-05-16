package main

import (
	"fmt"
	"strconv"
)

/*
874. 模拟行走机器人

机器人在一个无限大小的 XY 网格平面上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 commands ：
-2 ：向左转 90 度
-1 ：向右转 90 度
1 <= x <= 9 ：向前移动 x 个单位长度
在网格上有一些格子被视为障碍物 obstacles 。第 i 个障碍物位于网格点  obstacles[i] = (xi, yi) 。
机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。
返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）


注意：

北表示 +Y 方向。
东表示 +X 方向。
南表示 -Y 方向。
西表示 -X 方向。

示例 1：

输入：commands = [4,-1,3], obstacles = []
输出：25
解释：
机器人开始位于 (0, 0)：
1. 向北移动 4 个单位，到达 (0, 4)
2. 右转

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/walking-robot-simulation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func robotSim(commands []int, obstacles [][]int) int {
	ans := 0
	//记录障碍位置
	obs := make(map[string]bool)
	for _, o := range obstacles {
		obs[calcHash(o)] = true
	}

	//初始位置
	x := 0
	y := 0

	//初始方向
	dir := 0 // 北N=0, 东E=1, 南S=2, 西W=3，对应方向数组索引
	//方向数组
	var dx [4]int = [4]int{0, 1, 0, -1}
	var dy [4]int = [4]int{1, 0, -1, 0}

	for _, c := range commands {

		if c == -1 {
			//向左90度，对应方向变换
			dir = (dir + 1) % 4
		} else if c == -2 {
			//向右90度，对应方向变换
			dir = (dir + 3) % 4
		} else {
			for i := 0; i < c; i++ {
				nx := x + dx[dir]
				ny := y + dy[dir]
				//判断下一步是否是障碍物
				if _, ok := obs[calcHash([]int{nx, ny})]; ok {
					break
				}
				x = nx
				y = ny
				ans = Max(ans, (x*x + y*y))
			}
		}
	}
	return ans
}

//位置hash key
func calcHash(pos []int) string {
	return strconv.Itoa(pos[0]) + "," + strconv.Itoa(pos[1])
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {

	c := []int{4, -1, 3}
	obs := [][]int{}

	s := robotSim(c, obs)

	fmt.Println(s)

}
