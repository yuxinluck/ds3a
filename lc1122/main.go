package main

import "fmt"

func relativeSortArray(arr1 []int, arr2 []int) []int {

	//0 <= arr1[i], arr2[i] <= 1000，根据提示，最大值小于1000，可以创建长度为1001的数组
	var counter [1001]int
	for _, v := range arr1 {
		counter[v]++
	}

	ans := make([]int, 0, len(arr1))

	//遍历arr2,将counter中的arr2的值加入结果集，直到计数清零
	for _, v := range arr2 {
		for counter[v] > 0 {
			ans = append(ans, v)
			counter[v]--
		}
	}

	//将counter中剩余的值加入结果集，直到计数为零
	for i := 0; i < len(counter); i++ {
		for counter[i] > 0 {
			ans = append(ans, i)
			counter[i]--
		}
	}
	return ans
}

func main() {
	var arr1 []int = []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	var arr2 []int = []int{2, 1, 4, 3, 9, 6}

	fmt.Println(relativeSortArray(arr1, arr2))
}
