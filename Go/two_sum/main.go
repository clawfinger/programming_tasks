package main

import "fmt"

func twoSum(nums []int, target int) ([]int, bool) {
	cache := make(map[int]int)
	res := make([]int, 0)
	found := false
	for i, num := range nums {
		index, ok := cache[num]
		if ok {
			res = append(res, index, i)
			found = true
			break
		}
		cache[target-num] = i
	}
	return res, found
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		cache := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			_, ok := cache[-nums[i]-nums[j]]
			if ok {
				triplet := []int{nums[i], nums[j], -nums[i] - nums[j]}
				res = append(res, triplet)
			}
			cache[nums[j]] = 0
		}
	}
	return res
}

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -4})
	fmt.Println(res)
}
