package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			target := - nums[i]
			if nums[left]+nums[right] > target {
				right--
			} else if nums[left]+nums[right] < target{
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left - 1] {
					left++
				}
			}
		}
	}
	return res
}
