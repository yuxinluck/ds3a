package main

import (
	"fmt"
	"math/rand"
)



func sortArray(nums []int) []int {
    quickSort(nums,0, len(nums)-1)
	return nums
}

func quickSort(arr []int, l, r int) {
	
	if l >= r { return } 
	pivot := partition(arr,l,r)
	quickSort(arr,l,pivot)
	quickSort(arr,pivot+1,r)
}

func partition(arr []int, l,r int) int {
    pivot := l + rand.Intn(r - l +1)
	pivotVal := arr[pivot]

	for l <= r {
        for arr[l] < pivotVal { l++ }
		for arr[r] > pivotVal { r-- }

		if l == r { break }

		if l < r {
			arr[l] , arr[r] = arr[r],arr[l]
			l++
			r--
		}	
	}
	return r
}


func main() {

	nums := []int{5,2,3,1}

	fmt.Println(sortArray(nums))

}