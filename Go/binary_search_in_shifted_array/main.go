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
	return search_rec(nums, target, 0, len(nums)-1)
}

func main() {
	res := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Println(res)
}
