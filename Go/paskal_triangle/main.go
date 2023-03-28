package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{{1}, {1, 1}}

	current := res[1]
	for i := 2; i < numRows; i++ {
		left := 0
		right := 0
		generated := []int{}

		for j := 0; j < len(current)+1; j++ {
			if j == len(current) {
				right = 0
			} else {
				right = current[j]
			}
			next := left + right
			generated = append(generated, next)
			left = right
		}
		res = append(res, generated)
		current = generated
	}
	return res[:numRows]
}

func main() {
	fmt.Println(generate(5))
}
