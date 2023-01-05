package lucksaving

import "sort"

func LuckBalance(k int32, contests [][]int32) int32 {
	// Write your code here
	var luck int32 = 0
	canLose := k
	sort.Slice(contests, func(i, j int) bool {
		left := contests[i]
		right := contests[j]
		if left[1] == right[1] {
			return left[0] > right[0]
		}
		return left[1] > right[1]
	})
	for _, pair := range contests {
		valid := pair[1] == 1
		if canLose == 0 {
			if valid {
				luck -= pair[0]
			} else {
				luck += pair[0]
			}
		} else {
			if valid {
				canLose--
			}
			luck += pair[0]
		}
	}
	return luck
}
