package main

import (
	"fmt"
	"sort"
)



func merge(intervals [][]int) [][]int {
    
	//先进行排序
	sort.Slice(intervals,  func (i,j int) bool {
        return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	} )
	ans := make([][]int,0)

	var farthest int = -1
	var start = -1

	//遍历intervals
	for _,v := range intervals {
		left := v[0]
		right := v[1]
		if left <= farthest {
			farthest = Max(farthest,right)
		} else {
			if farthest != -1 {
				ans = append(ans, []int{start,farthest})
			}
			start = left
			farthest = right
		}
	}
    ans = append(ans, []int{start,farthest})
	return ans
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

	var intervals [][]int = [][]int{ []int{2,6}, []int{1,3}, []int{8,10}, []int{15,18}}
    
	fmt.Println(merge(intervals))

}



