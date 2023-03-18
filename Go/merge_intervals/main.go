package main

import (
	"fmt"
	"sort"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(left int, right int) bool { return intervals[left][0] < intervals[right][0] })
	start := intervals[0][0]
	res := make([][]int, 0)
	possibleEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if possibleEnd < intervals[i][0] {
			res = append(res, []int{start, possibleEnd})
			start = intervals[i][0]
		}
		possibleEnd = Max(possibleEnd, intervals[i][1])

	}
	res = append(res, []int{start, possibleEnd})
	return res
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 3}, {5, 7}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}
