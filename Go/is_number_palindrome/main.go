package main

import "math"

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	div := x / 10
	meaning := x % 10
	res := []int{meaning}

	for div != 0 {
		meaning = div % 10
		div = div / 10
		res = append(res, meaning)
	}
	tmp := 0
	for i, num := range res {
		tens := int(math.Pow10(len(res) - i - 1))
		tmp += num * tens
	}
	return tmp == x
}

func main() {
	isPalindrome(123)
}
