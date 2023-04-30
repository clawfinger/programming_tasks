package main

func removeElement(nums []int, val int) int {
	right := len(nums) - 1
	i := 0
	for i < right {
		if nums[i] == val {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			continue
		}
		i++
	}
	return i + 1
}
