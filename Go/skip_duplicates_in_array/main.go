package main

import "fmt"

func skip(arr []int) {
	for i := 0; i < len(arr); i++ {
		if i != 0 && arr[i] == arr[i-1] {
			continue
		}
		fmt.Print(arr[i], " ")
	}
}

func main() {
	skip([]int{2, 2, 2, 2, 2, 7})
}
