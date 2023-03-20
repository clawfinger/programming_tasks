package main

import "fmt"

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	length := 0
	for _, num := range nums {
		_, ok := set[num-1]
		if !ok {
			currentLength := 1
			for {
				_, ok := set[num+currentLength]
				if ok {
					currentLength++
				} else {
					length = max(length, currentLength)
					break
				}
			}
		}
	}
	return length
}

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 5, 7}))
}
