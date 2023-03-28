package main

import "fmt"

func moveZeroes(nums []int) {
	l := 0

	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	// arr := []int{0, 1, 2, 3, 12}
	// arr := []int{0, 0, 0, 0, 12}
	// arr := []int{1, 0, 0}
	moveZeroes(arr)
	fmt.Println(arr)
}
