package main

func findDisappearedNumbers(nums []int) []int {
    cache := make(map[int]struct{})
	for _, num := range nums {
		cache[num] = struct{}{}
	}
	ret := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		_, ok := cache[i]
		if !ok {
			ret = append(ret, i)
		}
	}
	return ret
}