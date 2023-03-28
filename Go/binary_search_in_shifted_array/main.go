package main

import "fmt"

func search_rec(nums []int, target int, i int, j int) int {
	mid := i + (j-i)/2
	if nums[mid] == target {
		return mid
	}
	left := -1
	right := -1
	if i <= mid-1 {
		left = search_rec(nums, target, i, mid-1)
	}
	if j >= mid+1 {
		right = search_rec(nums, target, mid+1, j)
	}
	if left == -1 && right == -1 {
		return -1
	} else if left == -1 {
		return right
	} else {
		return left
	}
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] < target { // going right
			if nums[right] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // going left
			if nums[left] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	res := search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)
	fmt.Println(res)
}
