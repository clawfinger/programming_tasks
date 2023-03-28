package main

import "fmt"

func removeDuplicates1(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		res = append(res, nums[i])
	}
	return res
}

func removeDuplicates(nums []int) int {
	paste := 1
	idx := 1
	for ; idx < len(nums); idx++ {
		if nums[idx-1] != nums[idx] {
			if paste != idx {
				nums[paste] = nums[idx]
			}
			paste++
		}
	}
	return paste
}

func main() {
	fmt.Println(removeDuplicates([]int{2, 7}))
}
