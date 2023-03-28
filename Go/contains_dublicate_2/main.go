package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]struct{})
	left := 0
	for i := 0; i < len(nums); i++ {
		if i >= k+1 {
			delete(window, nums[left])
			left++
		}
		_, ok := window[nums[i]]
		if ok {
			return true
		}

		window[nums[i]] = struct{}{}

	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 0))
}
