package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0

	lastPos := make(map[byte]int)

	res := 0

	for right < len(s) {
		last, ok := lastPos[s[right]]
		if ok {
			left = last + 1
		}
		res = max(res, right-left+1)
		lastPos[s[right]] = right
		right++
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("12324567"))
}
