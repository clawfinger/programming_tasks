package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	first := math.MinInt
	second := math.MinInt
	third := math.MinInt

	for _, num := range nums {
		possible := num
		for {
			if possible > first {
				if first == math.MinInt {
					first = possible
					break
				} else {
					possible, first = first, possible
				}
			}
			if possible > second {
				if second == math.MinInt {
					second = possible
					break
				} else {
					possible, second = second, possible
				}
			}
			if possible > third {
				if third == math.MinInt {
					third = possible
				} else {
					possible, third = third, possible
				}
			} else {
				break
			}
		}
	}
	if third != math.MinInt {
		return third
	} else {
		return first
	}
}

func main() {
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}
