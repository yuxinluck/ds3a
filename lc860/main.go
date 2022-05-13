package main

import "fmt"

var coins map[int]int

func lemonadeChange(bills []int) bool {
	coins = make(map[int]int)
	coins[5] = 0
	coins[10] = 0
	coins[20] = 0

	for _, bill := range bills {
		coins[bill]++
		if !exchange(bill - 5) {
			return false
		}
	}

	return true
}

func exchange(amount int) bool {
	for _, x := range []int{20, 10, 5} {
		for amount >= x && coins[x] > 0 {
			amount -= x
			coins[x]--
		}
	}
	return amount == 0
}

func main() {

	bills := []int{5, 5, 5, 10, 20}
	fmt.Println(lemonadeChange(bills))
}
