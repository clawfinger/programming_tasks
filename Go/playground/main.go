package main

import "fmt"

// 00111
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j := i
			for nums[j] == 0 {
				j++
				if j == len(nums) {
					return
				}
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func sortArrayByParity(nums []int) []int {
	return nil
}

func main() {
	arr := []int{0, 0, 0, 0, 0, 0}
	moveZeroes(arr)
	fmt.Println(arr)
}
