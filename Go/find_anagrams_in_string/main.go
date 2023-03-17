package main

import "fmt"

func isAnagram(s string, p string) bool {
	cacheS := make(map[rune]int)
	cacheP := make(map[rune]int)
	if len(s) != len(p) {
		return false
	}
	for _, letter := range s {
		cacheS[letter]++
	}
	for _, letter := range p {
		cacheP[letter]++
	}
	for letter, count := range cacheS {
		p, ok := cacheP[letter]
		if !ok || p != count {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	for i := 0; i <= len(s)-len(p); i++ {
		if isAnagram(s[i:i+len(p)], p) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	res := findAnagrams("a", "a")
	fmt.Println(res)
}
