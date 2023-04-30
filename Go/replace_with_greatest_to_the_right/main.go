package main

func replaceElements(arr []int) []int {
	max := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			max = arr[i]
			arr[i] = -1
		} else {
			current := arr[i]
			arr[i] = max
			if current > max {
				max = current
			}
		}
	}
	return arr
}

func main() {

}
