package main

import (
	"fmt"
	"math"
)

func findMinRecursive(nums []int, left int, right int) int {

	if left == right || nums[left] < nums[right] {
		return nums[left]
	}

	mid := left + (right-left)/2
	minLeft := findMinRecursive(nums, left, mid)
	minRight := findMinRecursive(nums, mid+1, right)
	return int(math.Min(float64(minLeft), float64(minRight)))
}

func findMin(nums []int) int {
	return findMinRecursive(nums, 0, len(nums)-1)
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
