package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// abcabcbb
func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	cache := make(map[byte]int) //char - position

	res := 0
	for right < len(s) {
		_, ok := cache[s[right]] //ok false
		for ok {
			delete(cache, s[left])
			left++                  // 1
			_, ok = cache[s[right]] //false
		}

		cache[s[right]] = right    // a = 0, b = 1, c = 2
		res = max(res, len(cache)) // res = 3

		right++ // right = 3
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
