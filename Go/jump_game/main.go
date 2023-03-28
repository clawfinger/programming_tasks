package main

import "fmt"

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func canJumpRec(nums []int, from int, memo map[int]bool) bool {
	res, ok := memo[from]
	if ok {
		return res
	}
	if from > len(nums)-1 {
		return true
	}
	res = false
	for i := nums[from]; i > 0; i-- {
		res = res || canJumpRec(nums, from+i, memo)
	}
	memo[from] = res
	return res
}

func canJump(nums []int) bool {
	memo := make(map[int]bool)
	return canJumpRec(nums, 0, memo)
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
