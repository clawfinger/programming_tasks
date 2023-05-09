package main

import (
	"fmt"
	"sort"
)

// 00111
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j := i
			for nums[j] == 0 {
				j++
				if j == len(nums) {
					return
				}
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func sortArrayByParity(nums []int) []int {
	odd := 0
	even := len(nums) - 1
	for odd < even {
		for odd < len(nums)-1 && nums[odd]%2 != 1 {
			odd++
		}
		if even > 0 && nums[even]%2 != 0 {
			even--
		}
		if odd < even {
			nums[odd], nums[even] = nums[even], nums[odd]
		}
	}
	return nums
}

func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	return left
}

func isOneDiff(s1 string, s2 string) bool {
	var longer *string
	var shorter *string

	if len(s1) > len(s2) {
		longer = &s1
		shorter = &s2
	} else {
		longer = &s2
		shorter = &s1
	}
	if len(*longer)-len(*shorter) > 1 {
		return false
	}
	hadDiff := false
	for i, _ := range *longer {
		if i >= len(*shorter) {
			return !hadDiff
		}
		if (*shorter)[i] != (*longer)[i] {
			if hadDiff {
				return false
			} else {
				hadDiff = true
			}
		}

	}
	return true
}

// monday = 0, sunday = 6 //jan okt
func countFridays(day int) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	count := 0
	for _, month := range months {
		if day == 6 {
			count++
		}
		offset := month % 7
		day = (day + offset) % 7
	}
	return count
}

func getComplexWords(words []string) []string {
	cache := make(map[string]struct{})
	for _, word := range words {
		cache[word] = struct{}{}
	}
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			current := words[i] + words[j]
			_, ok := cache[current]
			if ok {
				res = append(res, current)
			}
		}
	}
	return res
}

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
}

func main() {
	words := []string{"crea", "create", "open", "te", "teopen"}
	fmt.Println(getComplexWords(words))
}
