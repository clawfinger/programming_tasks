package main

import "fmt"

func findRow(matrix [][]int, target int, i int, j int) int {
	for i != j {
		mid := i + (j-i)/2

		if target >= matrix[mid][0] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	if i >= len(matrix) {
		if target <= matrix[len(matrix)-1][len(matrix[len(matrix)-1])-1] {
			return len(matrix) - 1
		}
		return -1
	}
	return i - 1
}

func searchMatrix(matrix [][]int, target int) bool {
	row := findRow(matrix, target, 0, len(matrix))
	if row < 0 {
		return false
	}
	i := 0
	j := len(matrix[row]) - 1
	for i <= j {
		mid := i + (j-i)/2
		if matrix[row][mid] == target {
			return true
		}
		if target >= matrix[row][mid] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}

func main() {
	// res := searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 10)
	res := searchMatrix([][]int{{1}, {3}}, 2)

	fmt.Println(res)
}
