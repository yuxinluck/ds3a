package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
http://www.amoscloud.com/?p=3561

给定一个射击比赛成绩单
  包含多个选手若干次射击的成绩分数
  请对每个选手按其最高三个分数之和进行降序排名
  输出降序排名后的选手id序列
  条件如下
    1. 一个选手可以有多个射击成绩的分数，且次序不固定
    2. 如果一个选手成绩少于3个，则认为选手的所有成绩无效，排名忽略该选手
    3. 如果选手的成绩之和相等，则相等的选手按照其id降序排列

   输入描述:
     输入第一行
         一个整数N
         表示该场比赛总共进行了N次射击
         产生N个成绩分数
         2<=N<=100

     输入第二行
       一个长度为N整数序列
       表示参与每次射击的选手id
       0<=id<=99

     输入第三行
        一个长度为N整数序列
        表示参与每次射击选手对应的成绩
        0<=成绩<=100

   输出描述:
      符合题设条件的降序排名后的选手ID序列
   示例一
      输入:
        13
        3,3,7,4,4,4,4,7,7,3,5,5,5
        53,80,68,24,39,76,66,16,100,55,53,80,55
      输出:
        5,3,7,4
      说明:
        该场射击比赛进行了13次
        参赛的选手为{3,4,5,7}
        3号选手成绩53,80,55 最高三个成绩的和为188
        4号选手成绩24,39,76,66  最高三个成绩的和为181
        5号选手成绩53,80,55  最高三个成绩的和为188
        7号选手成绩68,16,100  最高三个成绩的和为184
        比较各个选手最高3个成绩的和
        有3号=5号>7号>4号
        由于3号和5号成绩相等  且id 5>3
        所以输出5,3,7,4
*/

func main() {
	//第一步，获取输入
	var N int
	var s1 string
	var s2 string
	ids := make([]int, 0)
	scores := make([]int, 0)

	// fmt.Println("输入N:")
	// fmt.Scan(&N)
	// fmt.Println("输入ids:")
	// fmt.Scan(&s1)
	// fmt.Println("输入scores:")
	// fmt.Scan(&s2)

	N = 13
	s1 = "3,3,7,4,4,4,4,7,7,3,5,5,5"
	s2 = "53,80,68,24,39,76,66,16,100,55,53,80,55"

	fmt.Println(N)

	l1 := strings.Split(s1, ",")
	for i := 0; i < len(l1); i++ {
		id, _ := strconv.Atoi(l1[i])
		ids = append(ids, id)
	}
	fmt.Println(ids)
	l2 := strings.Split(s2, ",")
	for i := 0; i < len(l2); i++ {
		score, _ := strconv.Atoi(l2[i])
		scores = append(scores, score)
	}
	fmt.Println(scores)

	id_score_map := make(map[int][]int)

	for i := 0; i < N; i++ {
		if v, ok := id_score_map[ids[i]]; !ok {
			id_score_map[ids[i]] = []int{scores[i]}
		} else {
			id_score_map[ids[i]] = append(v, scores[i])
		}
	}
	fmt.Println(id_score_map)
    ans := make([][]int,0)
	for k, v := range id_score_map {
		//删除参赛次数小于3的id
		if len(v) < 3 {
			delete(id_score_map, k)
		}
		//每个ID分数排序取前三
		sort.Slice(v, func(i int, j int) bool { return v[i] > v[j] })
		id_score_map[k] = v[0:3]
        sum := 0
        for i:=0;i<3;i++{
			sum+=v[i]
		}
        ans = append(ans, []int{k,sum})	
	}
	fmt.Println(id_score_map)
	fmt.Println(ans)
	sort.Slice(ans,func(i, j int) bool {
		return ans[i][1] > ans[j][1] || (ans[i][1] == ans[j][1] && ans[i][0] > ans[j][0])
	})

	fmt.Println(ans)

	for i:=0 ;i< len(ans);i++ {
		fmt.Print(ans[i][0])
		
	}
	fmt.Println()
	
}
