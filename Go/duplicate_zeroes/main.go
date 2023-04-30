package main

//1,0,2,3,0,4,5,0
func duplicateZeros(arr []int)  {
	next := false
	tmp := -1
	for i := 0; i < len(arr); i++ {
		if next {
			tmp, arr[i] = arr[i], tmp
			arr[i] = 0
			next = false
		}
		if arr[i] == 0 {
			next = true
		} else {
			
		}
	}
}