package main

import (
	"fmt"
	"strconv"
	"strings"
)

var s = `1,0,0,0,1
0,0,0,1,1
0,1,0,1,0
1,0,0,1,1
1,0,1,0,1`

func main() {
	k := 5
	nums := make([][]string, 0)
	res := make([]int64, 0)
	for _, num := range strings.Split(s, "\n") {
		tmp := strings.Split(num, ",")
		nums = append(nums, tmp)
	}
	// fmt.Println(nums)
	for _, num := range nums {
		max, _ := strconv.ParseInt("0b"+strings.Join(num, ""), 0, 0)
		//strconv.Atoi 只能接受十进制格式字符串
		// max,_ := strconv.Atoi("0b"+strings.Join(num,""))

		//遍历每一种数字值，保留大值
		for i := 1; i < k; i++ {
			num = append(num[1:], num[0])
			tmp, _ := strconv.ParseInt("0b"+strings.Join(num, ""), 0, 0)
			if max < tmp {
				max = tmp
			}
		}
		res = append(res, max)
	}
    
	fmt.Println(res)

	ans := int64(0)

	for _,v := range  res{
        ans += v
	}
	fmt.Println(ans)
}

