package main

import "fmt"


func maxSubArray(nums []int) int {
	n := len(nums)
    f := make([]int,n)
	f[0] = nums[0]
    
	for i:=1;i<n;i++ {
		f[i] = Max(f[i-1]+nums[i],nums[i])
	}

	ans := f[0]

	for i:=1;i<n;i++ {
		ans = Max(ans,f[i])
	}
	return ans
}

func Max(x,y int) int {
	if x > y{
		return x
	}
	return y
}

func main() {
    nums := []int{-2,1,-3,4,-1,2,1,-5,4}

	fmt.Println(maxSubArray(nums))
}

