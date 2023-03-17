package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	cache := make(map[string][]string)
	for i := range strs {
		runes := []rune(strs[i])
		sort.Slice(runes, func(i int, j int) bool { return runes[i] < runes[j] })
		cache[string(runes)] = append(cache[string(runes)], strs[i])
	}
	res := make([][]string, 0)
	for _, words := range cache {
		res = append(res, words)
	}
	return res
}

func main() {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(res)
}
