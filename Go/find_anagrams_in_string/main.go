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

func findAnagramsBad(s string, p string) []int {
	res := make([]int, 0)
	for i := 0; i <= len(s)-len(p); i++ {
		if isAnagram(s[i:i+len(p)], p) {
			res = append(res, i)
		}
	}
	return res
}

func compareMaps(window map[byte]int, p map[byte]int) bool {
	for key, count := range window {
		if p[key] != count {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	cacheP := make(map[byte]int)
	cacheWindow := make(map[byte]int)
	for _, letter := range p {
		cacheP[byte(letter)]++
	}
	res := []int{}
	tail := 0
	forward := 0
	for forward < len(s) {

		if forward-tail >= len(p) {
			tailCount := cacheWindow[s[tail]]
			if tailCount == 1 {
				delete(cacheWindow, s[tail])
			} else {
				cacheWindow[s[tail]]--
			}
			tail++
		}
		cacheWindow[s[forward]]++
		if forward >= len(p)-1 {
			if compareMaps(cacheWindow, cacheP) {
				res = append(res, tail)
			}
		}
		forward++
	}
	return res
}

func main() {
	res := findAnagrams("abab", "ab")
	fmt.Println(res)
}
