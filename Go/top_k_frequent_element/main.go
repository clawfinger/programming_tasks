package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}
	type freq struct {
		num   int
		entry int
	}
	toSort := make([]freq, 0)
	for entry, count := range cache {
		toSort = append(toSort, freq{
			num:   count,
			entry: entry,
		})
	}
	sort.Slice(toSort, func(i, j int) bool {
		return toSort[i].num > toSort[j].num
	})
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, toSort[i].entry)
	}
	return res
}
