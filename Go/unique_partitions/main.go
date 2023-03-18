package main

import "fmt"

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}

func partitionLabels(s string) []int {
	cache := make(map[rune]int)
	for i, letter := range s {
		cache[letter] = i
	}
	end := 0
	size := 0
	res := make([]int, 0)
	for i, letter := range s {
		lastPos := cache[letter]
		end = max(end, lastPos)
		size++
		if i == end {
			res = append(res, size)
			size = 0
		}
	}
	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
