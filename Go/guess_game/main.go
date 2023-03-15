package main

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	i := 0
	j := n
	for {
		mid := i + (j-i)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res < 0 {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
}
