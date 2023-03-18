package main

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	water := 0
	for left < right {
		water = max(water, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return water
}
