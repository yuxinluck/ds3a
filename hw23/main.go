package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var s = `1T
20M
3G
10G6T
3M12G9M`

func main() {
	src := strings.Split(s, "\n")

	fmt.Println(src)

	//记录每个值的索引，与值
	res := make([][]int, 0)

	for i := 0; i < 5; i++ {
		num := src[i]
		sum := parseStr(num)
		fmt.Println(sum)
		res = append(res, []int{sum, i})
	}

	fmt.Println(res)

	//从小到大排序
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] < res[j][0] {
			return true
		}
		return false
	})

	//按照原格式输出
	for _, v := range res {
		fmt.Println(src[v[1]])
	}
}

func parseStr(s string) int {
	re1, _ := regexp.Compile("[MGT]")
	re2, _ := regexp.Compile("\\d+")

	l1 := strings.Split(strings.TrimSpace(re1.ReplaceAllString(s, " ")), " ")
	l2 := strings.Split(strings.TrimSpace(re2.ReplaceAllString(s, " ")), " ")

	// fmt.Println(l1)
	// fmt.Println(l2)
	lenth := len(l1)

	sum := 0

	for i := 0; i < lenth; i++ {
		n, _ := strconv.Atoi(l1[i])
		switch l2[i] {
		case "M":
			sum += n
		case "G":
			sum += n * 1024
		case "T":
			sum += n * 1024 * 1024
		}
	}

	return sum
}

/*
磁盘的容量单位常用的有
        M G T
        他们之间的换算关系为 1T =1024G 1G=1024M
        现在给定n块磁盘的容量，请对他们按从小到大的顺序进行稳定排序
        例如给定5块盘的容量
        5
        1T
        20M
        3G
        10G6T
        3M12G9M
        排序后的结果为
        20M
        3G
        3M 12G 9M
        1T,10G 6T
        注意单位可以重复出现
        上述3M 12G 9M表示的容量即为 3M 12G 9M 和12M 12G相等
        输入描述、
          输入第一行包含一个整数n
          2<=n<=100 表示磁盘的个数
          接下来的n行
          每行一个字符串
          2<长度<30
          表示磁盘的容量
          由一个或多个格式为MV的子串组成
          其中m表示容量大小
          v表示容量单位
          例如20M 1T
          磁盘容量的范围1~1024的正整数
          单位 M G T
         输出n行
         表示n块磁盘容量排序后的结果

         实例
         输入
         3
         1G
         2G
         1024M
        输出
        1G
        1024M
        2G
        说明：稳定排序要求相等值保留原来位置

        示例2
        3
        2G4m
        3M2G
        1T
        输出
        3M2G
        2G4M
        1T
*/
