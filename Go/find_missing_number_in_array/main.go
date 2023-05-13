package main

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		j := i
		number := nums[j]
		for number != j+1 {
			j = number - 1
			number, nums[j] = nums[j], number
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func findDisappearedNumbersAlternative(nums []int) []int {
	for _, num := range nums {
		index := abs(num) - 1
		nums[index] = abs(nums[index]) * -1
	}
	res := make([]int, 0)
	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// nums := []int{1, 2, 3, 4}

	findDisappearedNumbersAlternative(nums)
}
