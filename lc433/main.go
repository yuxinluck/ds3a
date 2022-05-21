package main

import "strings"

/*
433. 最小基因变化
https://leetcode.cn/problems/minimum-genetic-mutation/
*/

func minMutation(start string, end string, bank []string) int {

	//基因字母
	gene := []string{"A", "C", "G", "T"}
	depth := make(map[string]int)
	depth[start] = 0

	hashBank := make(map[string]int)

	for _, v := range bank {
		hashBank[v] = 1
	}

	if _, ok := hashBank[end]; !ok {
		return -1
	}

	//BFS
	q := make([]string, 0)
	q = append(q, start)

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		for i := 0; i < 8; i++ {
			for j := 0; j < 4; j++ {
				if strings.Split(s, "")[i] != gene[j] {
					temp := strings.Split(s, "")
					temp[i] = gene[j]
					ns := strings.Join(temp, "")
					if _, ok := hashBank[ns]; !ok {
						continue
					}
					if _, ok := depth[ns]; ok {
						continue
					}
					depth[ns] = depth[s] + 1
					q = append(q, ns)
					if ns == end {
						return depth[ns]
					}
				}
			}
		}
	}

	return -1

}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
