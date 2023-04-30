package main

//-4,-1,0,3,10
func abs(num int) int {
	if num >= 0 {
		return num
	}
	return num * -1
}

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, len(nums))
	idx := len(nums) - 1
	for left <= right {
		if abs(nums[left]) > abs(nums[right]) {
			res[idx] = nums[left] * nums[left]
			idx--
			left++
		} else {
			res[idx] = nums[right] * nums[right]
			idx--
			right--
		}
	}
	return res
}
