package main

import "fmt"

func jump(nums []int) int {

	now := 0
	ans := 0

	for now < len(nums)-1 {
		right := now + nums[now]
		if right >= len(nums)-1 {return ans+1}

		nextRight := right
		next := now

		for i := now+1 ; i <= right; i++ {
			if i+nums[i] > nextRight {
				nextRight = i+nums[i]
				next = i
			}
		} 

		now = next
		ans++
	}

	return ans
}

func main() {

	nums := []int{2,3,1,1,4}

	fmt.Println(jump(nums))

}