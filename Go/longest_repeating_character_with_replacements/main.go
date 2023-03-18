package main

import "fmt"

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func getMax(m map[byte]int) int {
	max := 0
	for _, value := range m {
		if max < value {
			max = value
		}
	}
	return max
}

func characterReplacement(s string, k int) int {
	left := 0
	right := 0
	res := 0
	cache := make(map[byte]int)
	for right < len(s) {
		cache[s[right]]++
		maxFrequent := getMax(cache)
		windowLen := right - left + 1
		if k >= windowLen-maxFrequent {
			res = max(res, windowLen)
		} else {
			cache[s[left]]--
			left++
		}
		right++
	}
	return res
}

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}
