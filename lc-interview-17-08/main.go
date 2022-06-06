package main

/*
https://leetcode.cn/problems/circus-tower-lcci/
面试题 17.08. 马戏团人塔

有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。
出于实际和美观的考虑，在上面的人要比下面的人矮一点且轻一点。已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/circus-tower-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import "sort"

type Person struct {
	height int
	weight int
}

func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	persons := make([]Person, n)
	for i := range persons {
		persons[i].height = height[i]
		persons[i].weight = weight[i]
	}
	sort.Slice(persons, func(i, j int) bool {
		// 身高高的在前边，身高相等则体重轻的在前边
		if persons[i].height == persons[j].height {
			return persons[i].weight < persons[j].weight
		}
		return persons[i].height > persons[j].height
	})
	var result []Person
	for _, p := range persons {
		// 在结果中找到第一个p不能叠在上面的人, 二分法
		j := sort.Search(len(result), func(i int) bool {
			c := result[i]
			return c.height <= p.height || c.weight <= p.weight
		})
		// 将第j个人替换成p
		if j == len(result) {
			result = append(result, p)
		} else {
			result[j] = p
		}
	}
	return len(result)
}
