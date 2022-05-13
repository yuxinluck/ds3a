package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := 1
	if len(nums) <= 1 {
		return len(nums)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	fmt.Println(nums[:n])
	return n
}

func main() {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))
}
