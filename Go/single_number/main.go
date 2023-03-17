package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, number := range nums {
		res = res ^ number
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 4, 4}))
}
