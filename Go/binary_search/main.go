package main

import "fmt"

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}

func main() {
	res := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(res)
}
