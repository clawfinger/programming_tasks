package queuebribes

import (
	"fmt"
)

func MinimumBribesBubbleSort(q []int32) {
	// Write your code here
	total := 0
	cache := make(map[int32]int)
	// bubble sort
	for i := 0; i < len(q); i++ {
		for j := i; j < len(q); j++ {
			left := q[i]
			right := q[j]
			if left > right {
				total++
				cache[q[i]]++
				q[i], q[j] = q[j], q[i]
			}
		}
	}
	for _, swaps := range cache {
		if swaps > 2 {
			fmt.Println("Too chaotic")
			return
		}
	}
	fmt.Println(total)
}

func MinimumBribes(q []int32) {
	total := 1
	cache := make(map[int32]int)
	// bubble sort
	currentMisplacedPos := int32(0)
	for i := 0; i < len(q); i++ {
		expected := int32(i + 1)
		if q[i] != expected {
			currentMisplacedPos = int32(i)
			break
		}
	}
	currentMisplaced := q[currentMisplacedPos]

	for {
		if (currentMisplaced - (currentMisplacedPos + 1)) > 2 {
			fmt.Printf("Too chaotic")
			return
		}
		_, ok := cache[currentMisplaced]
		if ok {
			fmt.Println(total)
			return
		}
		// picking misplaced array value currentMisplaced
		// calculation corrent index for this
		cache[currentMisplaced]++
		correctPos := currentMisplaced - 1
		// incrementing total
		total++
		// swaping current misplaced with value taking wrong position
		currentMisplaced = q[correctPos]
		currentMisplacedPos = correctPos
	}
}
