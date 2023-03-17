package main

func search_rec(nums []int, target int, i int, j int) bool {
	mid := i + (j-i)/2
	if nums[mid] == target {
		return true
	}
	left := false
	right := false
	if i <= mid-1 {
		left = search_rec(nums, target, i, mid-1)
	}
	if j >= mid+1 {
		right = search_rec(nums, target, mid+1, j)
	}
	return left || right
}

func search(nums []int, target int) bool {
	return search_rec(nums, target, 0, len(nums)-1)
}

func main() {
	search([]int{1, 0, 1, 1, 1}, 0)
}
