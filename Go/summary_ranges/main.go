package main

import (
	"fmt"
	"strconv"
	"strings"
)

func summaryRanges(nums []int) []string {
	left := 0
	right := 0

	res := make([]string, 0)
	var sb strings.Builder

	for right < len(nums) {
		if right == len(nums)-1 || nums[right+1]-nums[right] != 1 {
			sb.WriteString(strconv.Itoa(nums[left]))

			if right-left >= 1 {
				sb.WriteString("->")
				sb.WriteString(strconv.Itoa(nums[right]))
			}
			res = append(res, sb.String())
			sb.Reset()
			right++
			left = right
		} else {
			right++
		}
	}
	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
