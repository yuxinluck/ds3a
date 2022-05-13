package main

import "fmt"


func maxProduct(nums []int) int {
	n := len(nums)
    fMax := make([]int,n)
	fMin := make([]int,n)
	fMax[0] = nums[0]
	fMin[0] = nums[0]
    
	for i:=1;i<n;i++ {
		fMax[i] = Max(Max(fMax[i-1]*nums[i],nums[i]),fMin[i-1]*nums[i])
		fMin[i] = Min(Min(fMax[i-1]*nums[i],nums[i]),fMax[i-1]*nums[i])
	}

	ans := fMax[0]

	for i:=1;i<n;i++ {
		ans = Max(ans,fMax[i])
	}
	return ans
}

func Max(x,y int) int {
	if x > y{
		return x
	}
	return y
}

func Min(x,y int) int {
	if x > y{
		return y
	}
	return x
}



func main() {
    nums := []int{2,3,-2,4}

	fmt.Println(maxProduct(nums))
}