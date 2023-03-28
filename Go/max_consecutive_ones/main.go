package main

import "fmt"

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func findMaxConsecutiveOnesBad(nums []int) int { // 1,1,1,0,0,1,1
	res := 0
	current := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			current++
		} else {
			res = max(res, current)
			current = 0
		}
	}
	res = max(res, current)

	return res
}

func findMaxConsecutiveOnes(nums []int) int { // 1,1,1,0,0,1,1
	left := 0
	right := 0
	res := 0
	for right < len(nums) {
		if nums[right] == 0 {
			left = right
		}
		windowLength := right - left
		res = max(res, windowLength)
		right++
	}
	return res
}

func main() {
	arr := []int{1, 1, 1, 0, 0, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(arr))
}

