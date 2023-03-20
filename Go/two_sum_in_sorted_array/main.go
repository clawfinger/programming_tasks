package main

import "fmt"

func binarySearch(numbers []int, target int, exclude int) int {
	left := 0
	right := len(numbers) - 1
	for left <= right {
		mid := left + (right-left)/2
		if numbers[mid] == target {
			if mid == exclude {
				if mid+1 < len(numbers) && numbers[mid+1] == target {
					return mid + 1
				}
				if mid-1 >= 0 && numbers[mid+1] == target {
					return mid - 1
				}
			}
			return mid
		}
		if numbers[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func twoSum(numbers []int, target int) []int {
	for i, num := range numbers {
		second := binarySearch(numbers, target-num, i)
		if second != -1 {
			return []int{i + 1, second + 1}
		}
	}
	return []int{}
}

func twoSumPointers(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		check := numbers[left] + numbers[right]
		if check == target {
			return []int{left + 1, right + 1}
		}
		if check > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}

func main() {
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumPointers([]int{1, 2, 3, 4, 4, 9, 56, 90}, 8))

}
