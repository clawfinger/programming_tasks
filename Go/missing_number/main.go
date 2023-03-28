package main

func missingNumber(nums []int) int {
	order := make([]bool, len(nums)+1)
	for _, num := range nums {
		order[num] = true
	}
	for i, present := range order {
		if !present {
			return i
		}
	}
	return -1
}
