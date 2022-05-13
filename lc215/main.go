package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	//求第k大，就是第n-k+1 小，下标为len(nums)-1-k+1 ，即len(nums)-k
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

//求排序后，下表在index的值
func quickSort(arr []int, l int, r int, index int) int {

	if l >= r {
		return arr[l]
	}
	pivot := partition(arr, l, r)
	if index <= pivot {
		return quickSort(arr, l, pivot, index)
	} else {
		return quickSort(arr, pivot+1, r, index)
	}
}

func partition(arr []int, l, r int) int {
	pivot := l + rand.Intn(r-l+1)
	pivotVal := arr[pivot]

	for l <= r {
		for arr[l] < pivotVal {
			l++
		}
		for arr[r] > pivotVal {
			r--
		}

		if l == r {
			break
		}

		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	return r
}

func main() {

	nums := []int{3, 2, 1, 5, 6, 4}

	fmt.Println(findKthLargest(nums, 4))

	// fmt.Println(sortArray(nums))

}
