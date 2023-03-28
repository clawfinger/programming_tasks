package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	counter := 1
	element := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == element {
			counter++
		} else {
			counter--
		}
		if counter == 0 {
			element = nums[i]
			counter++
		}
	}

	return element
}

func main() {
	fmt.Println(majorityElement([]int{1, 3, 1, 1, 4, 1, 1, 5, 1, 1, 6, 2, 2}))
}
