package main

func search(nums []int, target int, leftMost bool) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if leftMost {
				if mid == 0 || nums[mid-1] != target {
					return mid
				}
				right = mid - 1
			} else {
				if mid == len(nums)-1 || nums[mid+1] != target {
					return mid
				}
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	res[0] = search(nums, target, true)
	res[1] = search(nums, target, false)
	return res
}
